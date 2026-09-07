package auth

import (
	"context"

	"google.golang.org/grpc"
)

type AuthFunc func(ctx context.Context) (context.Context, error)

type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

func UnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(authFunc AuthFunc) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
