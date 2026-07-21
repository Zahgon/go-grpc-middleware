package validator

import (
	"context"
)

type options struct {
	shouldFailFast          bool
	onValidationErrCallback OnValidationErrCallback
}
type Option func(*options)

func evaluateOpts(opts []Option) *options { _ = "STUB: not implemented"; return nil }

type OnValidationErrCallback func(ctx context.Context, err error)

func WithOnValidationErrCallback(onValidationErrCallback OnValidationErrCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFailFast() Option { _ = "STUB: not implemented"; return *new(Option) }
