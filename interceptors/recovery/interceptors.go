package recovery

import (
	"context"

	"google.golang.org/grpc"
)

type RecoveryHandlerFunc func(p any) (err error)

type RecoveryHandlerFuncContext func(ctx context.Context, p any) (err error)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func recoverFrom(ctx context.Context, p any, r RecoveryHandlerFuncContext) error {
	_ = "STUB: not implemented"
	return nil
}

type PanicError struct {
	Panic any
	Stack []byte
}

func (e *PanicError) Error() string { _ = "STUB: not implemented"; return "" }
