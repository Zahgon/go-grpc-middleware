package logging

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"google.golang.org/grpc/codes"
)

type LoggableEvent uint

const (
	StartCall LoggableEvent = iota

	FinishCall

	PayloadReceived

	PayloadSent
)

func has(events []LoggableEvent, event LoggableEvent) bool { _ = "STUB: not implemented"; return false }

var defaultOptions = &options{
	loggableEvents:    []LoggableEvent{StartCall, FinishCall},
	codeFunc:          DefaultErrorToCode,
	durationFieldFunc: DefaultDurationToFields,

	levelFunc:            nil,
	timestampFormat:      time.RFC3339,
	disableGrpcLogFields: nil,
}

type options struct {
	levelFunc               CodeToLevel
	loggableEvents          []LoggableEvent
	errorToFieldsFunc       ErrorToFields
	codeFunc                ErrorToCode
	durationFieldFunc       DurationToFields
	timestampFormat         string
	fieldsFromCtxCallMetaFn fieldsFromCtxCallMetaFn
	disableGrpcLogFields    []string
}

type Option func(*options)

func evaluateServerOpt(opts []Option) *options { _ = "STUB: not implemented"; return nil }

func evaluateClientOpt(opts []Option) *options { _ = "STUB: not implemented"; return nil }

type DurationToFields func(duration time.Duration) Fields

type ErrorToFields func(err error) Fields

type ErrorToCode func(err error) codes.Code

func DefaultErrorToCode(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

type CodeToLevel func(code codes.Code) Level

func DefaultServerCodeToLevel(code codes.Code) Level { _ = "STUB: not implemented"; return *new(Level) }

func DefaultClientCodeToLevel(code codes.Code) Level { _ = "STUB: not implemented"; return *new(Level) }

type (
	fieldsFromCtxFn         func(ctx context.Context) Fields
	fieldsFromCtxCallMetaFn func(ctx context.Context, c interceptors.CallMeta) Fields
)

func WithFieldsFromContext(f fieldsFromCtxFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFieldsFromContextAndCallMeta(f fieldsFromCtxCallMetaFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLogOnEvents(events ...LoggableEvent) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithErrorFields(f ErrorToFields) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLevels(f CodeToLevel) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCodes(f ErrorToCode) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDurationField(f DurationToFields) Option { _ = "STUB: not implemented"; return *new(Option) }

var DefaultDurationToFields = DurationToTimeMillisFields

func DurationToTimeMillisFields(duration time.Duration) Fields {
	_ = "STUB: not implemented"
	return *new(Fields)
}

func DurationToDurationField(duration time.Duration) Fields {
	_ = "STUB: not implemented"
	return *new(Fields)
}

func durationToMilliseconds(duration time.Duration) float32 { _ = "STUB: not implemented"; return 0 }

func WithTimestampFormat(format string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDisableLoggingFields(disableGrpcLogFields ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
