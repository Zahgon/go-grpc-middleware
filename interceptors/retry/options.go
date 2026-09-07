package retry

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	DefaultRetriableCodes = []codes.Code{codes.ResourceExhausted, codes.Unavailable}

	defaultOptions = &options{
		max:            0,
		perCallTimeout: 0,
		includeHeader:  true,
		backoffFunc:    BackoffLinearWithJitter(50*time.Millisecond, 0.10),
		onRetryCallback: OnRetryCallback(func(ctx context.Context, attempt uint, err error) {
			logTrace(ctx, "grpc_retry attempt: %d, backoff for %v", attempt, err)
		}),
		retriableFunc: newRetriableFuncForCodes(DefaultRetriableCodes),
	}
)

type BackoffFunc func(ctx context.Context, attempt uint) time.Duration

type OnRetryCallback func(ctx context.Context, attempt uint, err error)

type RetriableFunc func(err error) bool

func Disable() CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

func WithMax(maxRetries uint) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

func WithBackoff(bf BackoffFunc) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

func WithOnRetryCallback(fn OnRetryCallback) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

func WithCodes(retryCodes ...codes.Code) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

func WithPerRetryTimeout(timeout time.Duration) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

func WithRetriable(retriableFunc RetriableFunc) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type options struct {
	max             uint
	perCallTimeout  time.Duration
	includeHeader   bool
	backoffFunc     BackoffFunc
	onRetryCallback OnRetryCallback
	retriableFunc   RetriableFunc
}

type CallOption struct {
	grpc.EmptyCallOption
	applyFunc func(opt *options)
}

func reuseOrNewWithCallOptions(opt *options, callOptions []CallOption) *options {
	_ = "STUB: not implemented"
	return nil
}

func filterCallOptions(callOptions []grpc.CallOption) (grpcOptions []grpc.CallOption, retryOptions []CallOption) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newRetriableFuncForCodes(codes []codes.Code) func(err error) bool {
	_ = "STUB: not implemented"
	return nil
}
