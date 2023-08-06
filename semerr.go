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

package semerr

import "errors"

// Code is the semerr representation of a gRPC canonical error code.
type Code uint32

// GRPCCode returns the gRPC code for err.
//
//   - If err is nil, returns OK and true.
//   - If err is a semerr, returns its code and true.
//   - If is not a semerr, returns Unknown and false.
func GRPCCode(err error) (c Code, ok bool) {
	if err == nil {
		return 0, true
	}

	var se interface{ GRPCCode() Code }
	if errors.As(err, &se) {
		return se.GRPCCode(), true
	}

	return UnknownError{}.GRPCCode(), false
}

// HTTPStatus returns the HTTP status code for err.
//
//   - If err is nil, returns 200 and true.
//   - If err is a semerr, returns its status true.
//   - If is not a semerr, returns 500 and false.
func HTTPStatus(err error) (status int, ok bool) {
	if err == nil {
		return 200, true
	}

	var se interface{ HTTPStatus() int }
	if errors.As(err, &se) {
		return se.HTTPStatus(), true
	}

	return UnknownError{}.HTTPStatus(), false
}

// FromGRPCCode returns the semerr corresponding to c.
//
//   - If c is OK or unmapped, returns err.
func FromGRPCCode(c Code, err error) error {
	fn, ok := errFromCode[c]
	if !ok {
		return err
	}
	return fn(err)
}

// FromHTTPStatus returns the semerr corresponding to status.
//
// Not all semantic errors map to distinct HTTP statuses. In case of conflicts,
// FromHTTPStatus tries to choose the better suited type for the status.
//
//   - If status is 200 or unmapped, returns err.
func FromHTTPStatus(status int, err error) error {
	fn, ok := errFromStatus[status]
	if !ok {
		return err
	}
	return fn(err)
}
