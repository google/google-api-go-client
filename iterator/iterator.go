// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package iterator provides support for standard Google API iterators.
// See https://github.com/GoogleCloudPlatform/gcloud-golang/wiki/Iterator-Guidelines.
package iterator

import (
	"errors"
	"fmt"
	"reflect"
)

// Done is returned by an iterator's Next method when the iteration is
// complete; when there are no more items to return.
var Done = errors.New("no more items in iterator")

// We don't support mixed calls to Next and NextPage because they play
// with the paging state in incompatible ways.
var errMixed = errors.New("iterator: Next and NextPage called on same iterator")

// PageInfo contains information about an iterator's paging state.
type PageInfo struct {
	// Token is the token used to retrieve the next page of items from the
	// API. You may set Token immediately after creating an iterator to
	// begin iteration at a particular point. If Token is the empty string,
	// the iterator will begin with the first eligible item.
	//
	// After the underlying API method is called to retrieve a page of items,
	// Token is set to the next-page token in the response.
	Token string

	// MaxSize is the maximum number of items returned by a call to the API.
	// Set MaxSize as a hint to optimize the buffering behavior of the iterator.
	// If zero, the page size is determined by the underlying service.
	//
	// Use Pager to retrieve a page of a specific, exact size.
	MaxSize int

	// The error state of the iterator. Manipulated by PageInfo.next and Pager.
	err error

	// Function that fetches a page from the underlying service. It should pass
	// the pageSize and pageToken arguments to the service, fill the buffer
	// with the results from the call, and return the next-page token returned
	// by the service.
	fetch func(pageSize int, pageToken string) (nextPageToken string, err error)

	// Function that clears the iterator's buffer, returning any currently buffered items.
	bufLen func() int

	// Function that returns the buffer, after setting the buffer variable to nil.
	takeBuf func() interface{}

	// Set to true on first call to PageInfo.next or Pager.NextPage. Used to check
	// for calls to both Next and NextPage with the same iterator.
	nextCalled, nextPageCalled bool
}

// NewPageInfo exposes internals for iterator implementations.
// It is not a stable interface.
var NewPageInfo = newPageInfo

// If an iterator can support paging, its iterator-creating method should call
// this (via the NewPageInfo variable above).
//
// The fetch, bufLen and takeBuf arguments provide access to the
// iterator's internal slice of buffered items. They behave as described in
// PageInfo, above.
//
// The return value is the PageInfo.next method bound to the returned PageInfo value.
// (Returning it avoids exporting PageInfo.next.)
func newPageInfo(fetch func(int, string) (string, error), bufLen func() int, takeBuf func() interface{}) (*PageInfo, func() (bool, error)) {
	pi := &PageInfo{
		fetch:   fetch,
		bufLen:  bufLen,
		takeBuf: takeBuf,
	}
	return pi, pi.next
}

// Remaining returns the number of items available before the iterator makes another API call.
func (pi *PageInfo) Remaining() int { return pi.bufLen() }

// next provides support for an iterator's Next function. The first return
// value reports whether the buffer has at least one element in it. The second
// return value is the error state of the iterator. It may be Done. Note that
// next may return (true, Done), in which case the iterator should continue
// returning items from the buffer. It should only return Done itself when next
// returns (false, Done).
func (pi *PageInfo) next() (bool, error) {
	pi.nextCalled = true
	if pi.nextPageCalled {
		pi.err = errMixed
		return false, pi.err
	}
	for pi.bufLen() == 0 {
		if pi.err != nil {
			return false, pi.err
		}
		if err := pi.fill(pi.MaxSize); err != nil {
			pi.err = err
			return false, pi.err
		}
		if pi.Token == "" {
			pi.err = Done
			return pi.bufLen() > 0, Done
		}
	}
	return true, pi.err
}

// Call the API to fill the buffer, using size and pi.Token. Set pi.Token to the
// next-page token returned by the call.
func (pi *PageInfo) fill(size int) error {
	tok, err := pi.fetch(size, pi.Token)
	if err != nil {
		pi.takeBuf() // clear the buffer
		return err
	}
	pi.Token = tok
	return nil
}

// Pageable is implemented by iterators that support paging.
type Pageable interface {
	// PageInfo returns paging information associated with the iterator.
	PageInfo() *PageInfo
}

// Pager supports retrieving iterator items a page at a time.
type Pager struct {
	pageInfo *PageInfo
	pageSize int
}

// NewPager returns a pager that uses iter. Calls to its NextPage method will
// obtain exactly pageSize items, unless fewer remain. The pageToken argument
// indicates where to start the iteration. Pass the empty string to start at
// the beginning, or pass a token retrieved from a call to Pager.NextPage.
//
// If you use an iterator with a Pager, you must not call Next on the iterator.
func NewPager(iter Pageable, pageSize int, pageToken string) *Pager {
	p := &Pager{
		pageInfo: iter.PageInfo(),
		pageSize: pageSize,
	}
	p.pageInfo.Token = pageToken
	if pageSize <= 0 {
		p.pageInfo.err = errors.New("iterator: page size must be positive")
	}
	return p
}

// NextPage retrieves a sequence of items from the iterator and appends them
// to slicep, which must be a pointer to a slice of the iterator's item type.
// Exactly p.pageSize items will be appended, unless fewer remain.
//
// The first return value is the page token to use for the next page of items.
// If empty, there are no more pages. Aside from checking for the end of the
// iteration, the returned page token is only needed if the iteration is to be
// resumed a later time, in another context (possibly another process).
//
// The second return value is non-nil if an error occurred. It will never be
// the special iterator sentinel value Done. To recognize the end of the
// iteration, compare nextPageToken to the empty string.
func (p *Pager) NextPage(slicep interface{}) (nextPageToken string, err error) {
	p.pageInfo.nextPageCalled = true
	if p.pageInfo.err != nil {
		return "", p.pageInfo.err
	}
	if p.pageInfo.nextCalled {
		p.pageInfo.err = errMixed
		return "", p.pageInfo.err
	}
	if p.pageInfo.bufLen() > 0 {
		return "", errors.New("must call NextPage with an empty buffer")
	}
	// The buffer must be empty here, so takeBuf is a no-op. We call it just to get
	// the buffer's type.
	wantSliceType := reflect.PtrTo(reflect.ValueOf(p.pageInfo.takeBuf()).Type())
	if slicep == nil {
		return "", errors.New("nil passed to Pager.NextPage")
	}
	vslicep := reflect.ValueOf(slicep)
	if vslicep.Type() != wantSliceType {
		return "", fmt.Errorf("slicep should be of type %s, got %T", wantSliceType, slicep)
	}
	for p.pageInfo.bufLen() < p.pageSize {
		if err := p.pageInfo.fill(p.pageSize - p.pageInfo.bufLen()); err != nil {
			p.pageInfo.err = err
			return "", p.pageInfo.err
		}
		if p.pageInfo.Token == "" {
			break
		}
	}
	e := vslicep.Elem()
	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.pageInfo.takeBuf())))
	return p.pageInfo.Token, nil
}
