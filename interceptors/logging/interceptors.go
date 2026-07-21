package logging

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"google.golang.org/grpc"
)

type reporter struct {
	interceptors.CallMeta

	ctx             context.Context
	kind            string
	startCallLogged bool

	opts   *options
	fields Fields
	logger Logger
}

func (c *reporter) PostCall(err error, duration time.Duration) { _ = "STUB: not implemented"; return }

func (c *reporter) PostMsgSend(payload any, err error, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *reporter) PostMsgReceive(payload any, err error, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func reportable(logger Logger, opts *options) interceptors.CommonReportableFunc {
	_ = "STUB: not implemented"
	return *new(interceptors.CommonReportableFunc)
}

func UnaryClientInterceptor(logger Logger, opts ...Option) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor(logger Logger, opts ...Option) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

func UnaryServerInterceptor(logger Logger, opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(logger Logger, opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
