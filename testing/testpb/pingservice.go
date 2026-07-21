package testpb

import (
	"context"
)

const (
	ListResponseCount = 100
)

var TestServiceFullName = TestService_ServiceDesc.ServiceName

var _ TestServiceServer = &TestPingService{}

type TestPingService struct {
	UnimplementedTestServiceServer
	PingFunc func(ctx context.Context)
}

func (s *TestPingService) PingEmpty(_ context.Context, _ *PingEmptyRequest) (*PingEmptyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TestPingService) Ping(ctx context.Context, ping *PingRequest) (*PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TestPingService) PingError(_ context.Context, ping *PingErrorRequest) (*PingErrorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TestPingService) PingList(ping *PingListRequest, stream TestService_PingListServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TestPingService) PingStream(stream TestService_PingStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TestPingService) PingClientStream(stream TestService_PingClientStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}
