package validator

import (
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type recvWrapper struct {
	*options
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }
