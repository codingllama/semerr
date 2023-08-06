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

/*
Package semerr provides semantic error wrappers. It is designed to be simple,
small, fast, integrated with the stdlib and to pull zero external dependencies.

All semantic errors are simple structs, meant to be used by themselves or
wrapping another error:

	// Standalone use:
	err1 := semerr.NotFoundError{}

	// Annotate existing error:
	err2 := semerr.InvalidArgumentError{
		fmt.Errorf("invalid user name %q: max allowed length is %d", name, maxNameLen)
	}

Errors can be type checked using errors.As:

	if ok := errors.As(err, &semerr.NotFoundError{}); ok {
		// handle not found
	}

Errors can be converted to their gRPC or HTTP counterparts using the [GRPCCode]
and [HTTPStatus] functions.

	// Example of semerr to gRPC status conversion.
	// Suitable for a gRPC server interceptor.
	c, _ := semerr.GRPCCode(err)
	statusErr := status.Error(codes.Code(c), err.Error())

Semantic errors are based on the gRPC [canonical error codes].

[canonical error codes]: https://pkg.go.dev/google.golang.org/grpc/codes#pkg-overview
*/
package semerr
