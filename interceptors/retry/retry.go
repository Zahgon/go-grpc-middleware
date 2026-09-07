package retry

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
)

const (
	AttemptMetadataKey = "x-retry-attempt"
)

func UnaryClientInterceptor(optFuncs ...CallOption) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor(optFuncs ...CallOption) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

type serverStreamingRetryingStream struct {
	grpc.ClientStream
	bufferedSends []any
	wasClosedSend bool
	parentCtx     context.Context
	callOpts      *options
	streamerCall  func(ctx context.Context) (grpc.ClientStream, error)
	mu            sync.RWMutex
}

func (s *serverStreamingRetryingStream) setStream(clientStream grpc.ClientStream) {
	_ = "STUB: not implemented"
	return
}

func (s *serverStreamingRetryingStream) getStream() grpc.ClientStream {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream)
}

func (s *serverStreamingRetryingStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStreamingRetryingStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (s *serverStreamingRetryingStream) Header() (grpcMetadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(grpcMetadata.MD), nil
}

func (s *serverStreamingRetryingStream) Trailer() grpcMetadata.MD {
	_ = "STUB: not implemented"
	return *new(grpcMetadata.MD)
}

func (s *serverStreamingRetryingStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStreamingRetryingStream) receiveMsgAndIndicateRetry(m any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *serverStreamingRetryingStream) reestablishStreamAndResendBuffer(callCtx context.Context) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func waitRetryBackoff(attempt uint, parentCtx context.Context, callOpts *options) error {
	_ = "STUB: not implemented"
	return nil
}

func isRetriable(err error, callOpts *options) bool { _ = "STUB: not implemented"; return false }

func isContextError(err error) bool { _ = "STUB: not implemented"; return false }

func perCallContext(parentCtx context.Context, callOpts *options, attempt uint) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func perStreamContext(parentCtx context.Context, callOpts *options, attempt uint) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func contextErrToGrpcErr(err error) error { _ = "STUB: not implemented"; return nil }

func logTrace(ctx context.Context, format string, a ...any) { _ = "STUB: not implemented"; return }
