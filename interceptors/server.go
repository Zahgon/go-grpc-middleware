package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(reportable ServerReportable) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(reportable ServerReportable) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type monitoredServerStream struct {
	grpc.ServerStream

	newCtx   context.Context
	reporter Reporter
}

func (s *monitoredServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *monitoredServerStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *monitoredServerStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }
