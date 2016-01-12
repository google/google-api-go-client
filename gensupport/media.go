// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"google.golang.org/api/googleapi"
)

const sniffBuffSize = 512

func newContentSniffer(r io.Reader) *contentSniffer {
	return &contentSniffer{r: r}
}

// contentSniffer wraps a Reader, and reports the content type determined by sniffing up to 512 bytes from the Reader.
type contentSniffer struct {
	r     io.Reader
	start []byte // buffer for the sniffed bytes.
	err   error  // set to any error encountered while reading bytes to be sniffed.

	ctype   string // set on first sniff.
	sniffed bool   // set to true on first sniff.
}

func (cs *contentSniffer) Read(p []byte) (n int, err error) {
	// Ensure that the content type is sniffed before any data is consumed from Reader.
	_, _ = cs.ContentType()

	if len(cs.start) > 0 {
		n := copy(p, cs.start)
		cs.start = cs.start[n:]
		return n, nil
	}

	// We may have read some bytes into start while sniffing, even if the read ended in an error.
	// We should first return those bytes, then the error.
	if cs.err != nil {
		return 0, cs.err
	}

	// Now we have handled all bytes that were buffered while sniffing.  Now just delegate to the underlying reader.
	return cs.r.Read(p)
}

// ContentType returns the sniffed content type, and whether the content type was succesfully sniffed.
func (cs *contentSniffer) ContentType() (string, bool) {
	if cs.sniffed {
		return cs.ctype, cs.ctype != ""
	}
	cs.sniffed = true
	// If ReadAll hits EOF, it returns err==nil.
	cs.start, cs.err = ioutil.ReadAll(io.LimitReader(cs.r, sniffBuffSize))

	// Don't try to detect the content type based on possibly incomplete data.
	if cs.err != nil {
		return "", false
	}

	cs.ctype = http.DetectContentType(cs.start)
	return cs.ctype, true
}

// DetermineContentType determines the content type of the supplied reader.
// If the content type is already known, it can be specified via ctype.
// Otherwise, the content of media will be sniffed to determine the content type.
// If media implements googleapi.ContentTyper (deprecated), this will be used
// instead of sniffing the content.
// After calling DetectContentType the caller must not perform further reads on
// media, but rather read from the Reader that is returned.
func DetermineContentType(media io.Reader, ctype string) (io.Reader, string) {
	// Note: callers could avoid calling DetectContentType if ctype != "",
	// but doing the check inside this function reduces the amount of
	// generated code.
	if ctype != "" {
		return media, ctype
	}

	// For backwards compatability, allow clients to set content
	// type by providing a ContentTyper for media.
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return media, typer.ContentType()
	}

	sniffer := newContentSniffer(media)
	if ctype, ok := sniffer.ContentType(); ok {
		return sniffer, ctype
	}
	// If content type could not be sniffed, reads from sniffer will eventually fail with an error.
	return sniffer, ""
}

type typeReader struct {
	io.Reader
	typ string
}

// multipartReader combines the contents of multiple readers to creat a multipart/related HTTP body.
// Close must be called if reads from the multipartReader are abandoned before reaching EOF.
type multipartReader struct {
	parts    []typeReader
	pr       *io.PipeReader
	pipeOpen bool
	ctype    string
}

func newMultipartReader(parts []typeReader) *multipartReader {
	mp := &multipartReader{parts: parts}
	// Start a goroutine to feed the pipe.
	var pw *io.PipeWriter
	mp.pr, pw = io.Pipe()
	mp.pipeOpen = true
	mpw := multipart.NewWriter(pw)
	mp.ctype = "multipart/related; boundary=" + mpw.Boundary()
	go func() {
		for _, part := range mp.parts {
			w, err := mpw.CreatePart(typeHeader(part.typ))
			if err != nil {
				mpw.Close()
				pw.CloseWithError(fmt.Errorf("googleapi: CreatePart failed: %v", err))
				return
			}
			_, err = io.Copy(w, part.Reader)
			if err != nil {
				mpw.Close()
				pw.CloseWithError(fmt.Errorf("googleapi: Copy failed: %v", err))
				return
			}
		}

		mpw.Close()
		pw.Close()
	}()
	return mp
}

func (mp *multipartReader) Read(data []byte) (n int, err error) {
	return mp.pr.Read(data)
}

func (mp *multipartReader) Close() error {
	if !mp.pipeOpen {
		return nil
	}
	mp.pipeOpen = false
	return mp.pr.Close()
}

// CombineBodyMedia combines a json body with media content to create a multipart/related HTTP body.
// It returns a ReadCloser containing the combined body, and the overall "multipart/related" content type, with random boundary.
//
// The caller must call Close on the returned ReadCloser if reads are abandoned before reaching EOF.
func CombineBodyMedia(media io.Reader, mediaType string, body io.Reader, bodyType string) (io.ReadCloser, string) {
	mp := newMultipartReader([]typeReader{
		{body, bodyType},
		{media, mediaType},
	})
	return mp, mp.ctype
}

func getMediaType(media io.Reader) (io.Reader, string) {
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return media, typer.ContentType()
	}

	sniffer := newContentSniffer(media)
	typ, ok := sniffer.ContentType()
	if !ok {
		// TODO(mcgreevy): Remove this default.  It maintains the semantics of the existing code,
		// but should not be relied on.
		typ = "application/octet-stream"
	}
	return sniffer, typ
}

// DetectMediaType detects and returns the content type of the provided media.
// If the type can not be determined, "application/octet-stream" is returned.
func DetectMediaType(media io.ReaderAt) string {
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return typer.ContentType()
	}

	typ := "application/octet-stream"
	buf := make([]byte, 1024)
	n, err := media.ReadAt(buf, 0)
	buf = buf[:n]
	if err == nil || err == io.EOF {
		typ = http.DetectContentType(buf)
	}
	return typ
}

func typeHeader(contentType string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", contentType)
	return h
}
