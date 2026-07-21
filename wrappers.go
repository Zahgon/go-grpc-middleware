package middleware

import (
	"context"

	"google.golang.org/grpc"
)

type WrappedServerStream struct {
	grpc.ServerStream

	WrappedContext context.Context
}

func (w *WrappedServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	_ = "STUB: not implemented"
	return nil
}
