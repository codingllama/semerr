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
	"fmt"

	"github.com/codingllama/semerr"
)

func Example_annotate() {
	var err error = semerr.NotFoundError{errors.New("user not found")}
	fmt.Println(err)

	// Verify error type.
	fmt.Println(errors.As(err, &semerr.NotFoundError{}))

	// Get code and status.
	code, _ := semerr.GRPCCode(err)
	status, _ := semerr.HTTPStatus(err)
	fmt.Println(code, status)

	// Output:
	// user not found
	// true
	// 5 404
}

func Example_standalone() {
	err := semerr.NotFoundError{} // that's it!

	fmt.Println(err.Error())
	fmt.Println(err.GRPCCode())
	fmt.Println(err.HTTPStatus())

	// Output:
	// not found
	// 5
	// 404
}

func Example_wrap() {
	err := fmt.Errorf("user 10 %w", semerr.NotFoundError{})
	fmt.Println(err)

	// Error is identified as a NotFoundError.
	fmt.Println(errors.Is(err, semerr.NotFoundError{}))

	// Get code and status.
	code, _ := semerr.GRPCCode(err)
	status, _ := semerr.HTTPStatus(err)
	fmt.Println(code, status)

	// Output:
	// user 10 not found
	// true
	// 5 404
}

func Example_embed() {
	type UserNotFoundError struct {
		semerr.NotFoundError
	}

	err := &UserNotFoundError{}
	err.Err = errors.New("user 10 not found")

	var nfe semerr.NotFoundError
	ok := errors.As(err, &nfe)
	fmt.Println(ok)
	fmt.Println(nfe)

	// Output:
	// true
	// user 10 not found
}

func ExampleGRPCCode() {
	// nil returns OK and true.
	code, ok := semerr.GRPCCode(nil)
	fmt.Println(code, ok)

	// Semantic errors return their code and true.
	code, ok = semerr.GRPCCode(semerr.NotFoundError{})
	fmt.Println(code, ok)

	// Unknown errors return Unknown and false.
	code, ok = semerr.GRPCCode(errors.New("not a semantic error"))
	fmt.Println(code, ok)

	// Output:
	// 0 true
	// 5 true
	// 2 false
}

func ExampleHTTPStatus() {
	// nil returns 200 and true.
	status, ok := semerr.HTTPStatus(nil)
	fmt.Println(status, ok)

	// Semantic errors return their status and true.
	status, ok = semerr.HTTPStatus(semerr.NotFoundError{})
	fmt.Println(status, ok)

	// Unknown errors return 500 and false.
	status, ok = semerr.HTTPStatus(errors.New("not a semantic error"))
	fmt.Println(status, ok)

	// Output:
	// 200 true
	// 404 true
	// 500 false
}

// codes is a mock for the grpc/codes package.
var codes = struct{ NotFound uint32 }{
	NotFound: uint32(semerr.NotFoundError{}.GRPCCode()),
}

func ExampleFromGRPCCode() {
	// err is wrapped in a semerr.NotFoundError.
	err := semerr.FromGRPCCode(semerr.Code(codes.NotFound), errors.New("user not found"))
	fmt.Println(err)
	fmt.Println(errors.As(err, &semerr.NotFoundError{})) // Assert type.

	// Output:
	// user not found
	// true
}

func ExampleFromHTTPStatus() {
	// err is wrapped in a semerr.NotFoundError.
	err := semerr.FromHTTPStatus(404, errors.New("user not found"))
	fmt.Println(err)
	fmt.Println(errors.As(err, &semerr.NotFoundError{})) // Assert type.

	// Output:
	// user not found
	// true
}
