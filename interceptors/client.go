package interceptors

import (
	"time"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(reportable ClientReportable) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor(reportable ClientReportable) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

type monitoredClientStream struct {
	grpc.ClientStream

	startTime       time.Time
	hasServerStream bool
	reporter        Reporter
}

func (s *monitoredClientStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *monitoredClientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }
