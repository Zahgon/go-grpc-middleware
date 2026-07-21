package selector

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"google.golang.org/grpc"
)

type Matcher interface {
	Match(ctx context.Context, callMeta interceptors.CallMeta) bool
}

func MatchFunc(f func(ctx context.Context, callMeta interceptors.CallMeta) bool) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

type funcSelector struct {
	f func(ctx context.Context, callMeta interceptors.CallMeta) bool
}

func (s funcSelector) Match(ctx context.Context, callMeta interceptors.CallMeta) bool {
	_ = "STUB: not implemented"
	return false
}

func UnaryServerInterceptor(i grpc.UnaryServerInterceptor, matcher Matcher) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(i grpc.StreamServerInterceptor, matcher Matcher) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func UnaryClientInterceptor(i grpc.UnaryClientInterceptor, matcher Matcher) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor(i grpc.StreamClientInterceptor, matcher Matcher) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}
