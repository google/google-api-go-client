package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"golang.org/x/net/context"
	storage "google.golang.org/api/storage/v1"
)

const fname = "storage_test.go"

// RewriteTransport is an http.RoundTripper that rewrites requests
// using the provided URL's Scheme and Host, and its Path as a prefix.
// The Opaque field is untouched.
// If Transport is nil, http.DefaultTransport is used
type RewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// note that url.URL.ResolveReference doesn't work here
	// since t.u is an absolute url
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}

type myHandler struct {
	location string
	r        *http.Request
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r = r
	if h.location != "" {
		w.Header().Set("Location", h.location)
	}
	fmt.Fprintf(w, "{}")
}

func TestMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("unable to parse server URL %q: %v", server.URL, err)
	}
	client := &http.Client{
		Transport: RewriteTransport{URL: u},
	}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}

	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("unable to open test file: %v", err)
	}
	defer f.Close()

	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Media(f).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w := "POST"; g.Method != w {
		t.Errorf("unknown Method, got %q, want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("unknown Proto, got %q, want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("unknown ProtoMajor, got %v, want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("unknown ProtoMinor, got %v, want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w, k := "multipart/related; boundary=", "Content-Type"; len(g.Header[k]) != 1 || !strings.HasPrefix(g.Header[k][0], w) {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w := int64(-1); g.ContentLength != w {
		t.Errorf("unknown ContentLength, got %v, want %v", g.ContentLength, w)
	}
	if w := "chunked"; len(g.TransferEncoding) != 1 || g.TransferEncoding[0] != w {
		t.Errorf("unknown TransferEncoding, got %#v, want %q", g.TransferEncoding, w)
	}
	if w := "www.googleapis.com"; g.Host != w {
		t.Errorf("unknown Host, got %q, want %q", g.Host, w)
	}
	if g.Form != nil {
		t.Errorf("unknown Form, got %#v, want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("unknown PostForm, got %#v, want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("unknown MultipartForm, got %#v, want nil", g.MultipartForm)
	}
	if w := "http://www.googleapis.com/upload/storage/v1/b/mybucket/o?alt=json&uploadType=multipart"; g.RequestURI != w {
		t.Errorf("unknown RequestURI, got %q, want %q", g.RequestURI, w)
	}
}

func TestResumableMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("unable to parse server URL %q: %v", server.URL, err)
	}
	handler.location = u.String()
	client := &http.Client{
		Transport: RewriteTransport{URL: u},
	}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}

	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("unable to open test file: %v", err)
	}
	defer f.Close()
	fi, err := os.Stat(fname)
	if err != nil {
		t.Fatalf("unable to stat test file: %v", err)
	}

	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Name(fname).ResumableMedia(context.Background(), f, fi.Size(), "text/plain").Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w := "POST"; g.Method != w {
		t.Errorf("unknown Method, got %q, want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("unknown Proto, got %q, want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("unknown ProtoMajor, got %v, want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("unknown ProtoMinor, got %v, want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if k := "Content-Type"; len(g.Header[k]) != 0 {
		t.Errorf("unknown header %q, got %#v, want nil", k, g.Header[k])
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w := int64(0); g.ContentLength != w {
		t.Errorf("unknown ContentLength, got %v, want %v", g.ContentLength, w)
	}
	if len(g.TransferEncoding) != 0 {
		t.Errorf("unknown TransferEncoding, got %#v, want nil", g.TransferEncoding)
	}
	if g.Form != nil {
		t.Errorf("unknown Form, got %#v, want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("unknown PostForm, got %#v, want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("unknown MultipartForm, got %#v, want nil", g.MultipartForm)
	}
}

func TestNoMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("unable to parse server URL %q: %v", server.URL, err)
	}
	client := &http.Client{
		Transport: RewriteTransport{URL: u},
	}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}

	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w := "POST"; g.Method != w {
		t.Errorf("unknown Method, got %q, want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("unknown Proto, got %q, want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("unknown ProtoMajor, got %v, want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("unknown ProtoMinor, got %v, want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w, k := "application/json", "Content-Type"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("unknown header %q, got %#v, want %q", k, g.Header[k], w)
	}
	if w := int64(116); g.ContentLength != w {
		t.Errorf("unknown ContentLength, got %v, want %v", g.ContentLength, w)
	}
	if len(g.TransferEncoding) != 0 {
		t.Errorf("unknown TransferEncoding, got %#v, want []string{}", g.TransferEncoding)
	}
	if w := "www.googleapis.com"; g.Host != w {
		t.Errorf("unknown Host, got %q, want %q", g.Host, w)
	}
	if g.Form != nil {
		t.Errorf("unknown Form, got %#v, want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("unknown PostForm, got %#v, want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("unknown MultipartForm, got %#v, want nil", g.MultipartForm)
	}
	if w := "http://www.googleapis.com/storage/v1/b/mybucket/o?alt=json"; g.RequestURI != w {
		t.Errorf("unknown RequestURI, got %q, want %q", g.RequestURI, w)
	}
}
