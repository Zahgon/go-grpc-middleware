package protovalidate

import (
	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(validator protovalidate.Validator, opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(validator protovalidate.Validator, opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type wrappedServerStream struct {
	grpc.ServerStream

	validator protovalidate.Validator
	options   *options
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMsg(m interface{}, validator protovalidate.Validator, opts *options) error {
	_ = "STUB: not implemented"
	return nil
}
