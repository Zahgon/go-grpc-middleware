package ratelimit

import (
	"context"

	"google.golang.org/grpc"
)

type Limiter interface {
	Limit(ctx context.Context) error
}

func UnaryServerInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(limiter Limiter) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func UnaryClientInterceptor(limiter Limiter) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor(limiter Limiter) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}
