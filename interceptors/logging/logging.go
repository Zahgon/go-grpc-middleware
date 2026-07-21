package logging

import (
	"context"

	_ "log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
)

var (
	SystemTag = []string{"protocol", "grpc"}

	ComponentFieldKey    = "grpc.component"
	KindServerFieldValue = "server"
	KindClientFieldValue = "client"
	ServiceFieldKey      = "grpc.service"
	MethodFieldKey       = "grpc.method"
	MethodTypeFieldKey   = "grpc.method_type"
)

type (
	fieldsCtxMarker struct{}
	fieldsCtxValue  struct {
		fields Fields
	}
)

var fieldsCtxMarkerKey = &fieldsCtxMarker{}

func newCommonFields(kind string, c interceptors.CallMeta) Fields {
	_ = "STUB: not implemented"
	return *new(Fields)
}

func disableCommonLoggingFields(kind string, c interceptors.CallMeta, disableFields []string) Fields {
	_ = "STUB: not implemented"
	return *new(Fields)
}

type Fields []any

func (f Fields) Iterator() *iter { _ = "STUB: not implemented"; return nil }

type iter struct {
	f Fields
	i int
}

func (i *iter) Next() bool { _ = "STUB: not implemented"; return false }

func (i *iter) At() (k string, v any) { _ = "STUB: not implemented"; return "", *new(any) }

func (f *Fields) Delete(key string) { _ = "STUB: not implemented"; return }

func (f Fields) WithUnique(add Fields) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func (f Fields) AppendUnique(add Fields) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func ExtractFields(ctx context.Context) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func InjectFields(ctx context.Context, f Fields) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func InjectLogField(ctx context.Context, key string, val any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func AddFields(ctx context.Context, f Fields) { _ = "STUB: not implemented"; return }

type Logger interface {
	Log(ctx context.Context, level Level, msg string, fields ...any)
}

type LoggerFunc func(ctx context.Context, level Level, msg string, fields ...any)

func (f LoggerFunc) Log(ctx context.Context, level Level, msg string, fields ...any) {
	_ = "STUB: not implemented"
	return
}
