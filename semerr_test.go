// Copyright (c) 2023 Alan Parra
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package semerr_test

import (
	"errors"
	"testing"

	"github.com/codingllama/semerr"
)

func TestFrom_unmapped(t *testing.T) {
	err1 := errors.New("an error")

	fromCode := func(n interface{}, err error) error {
		c := semerr.Code(n.(int))
		return semerr.FromGRPCCode(c, err)
	}
	fromHTTP := func(s interface{}, err error) error {
		return semerr.FromHTTPStatus(s.(int), err)
	}

	tests := []struct {
		name string
		fn   func(interface{}, error) error
		in   interface{}
	}{
		{
			name: "FromGRPCCode(OK, $err)",
			fn:   fromCode,
			in:   0, // OK=0
		},
		{
			name: "FromGRPCCode(999, $err)",
			fn:   fromCode,
			in:   999, // 999 represents an unmapped code
		},
		{
			name: "FromHTTPStatus(200, $err)",
			fn:   fromHTTP,
			in:   200,
		},
		{
			name: "FromHTTPStatus(999, $err)",
			fn:   fromHTTP,
			in:   999, // 999 represents an unmapped status
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.fn(test.in, nil); got != nil {
				t.Errorf("fn(in, nil) = %v, want nil", got)
			}

			//nolint:errorlint // We really mean !=.
			if got := test.fn(test.in, err1); got != err1 {
				t.Errorf("fn(in, %q) = %#v, want %#v", err1, got, err1)
			}
		})
	}
}
