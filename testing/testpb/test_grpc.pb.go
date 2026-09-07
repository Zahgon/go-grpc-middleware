package testpb

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	TestService_PingEmpty_FullMethodName        = "/testing.testpb.v1.TestService/PingEmpty"
	TestService_Ping_FullMethodName             = "/testing.testpb.v1.TestService/Ping"
	TestService_PingError_FullMethodName        = "/testing.testpb.v1.TestService/PingError"
	TestService_PingList_FullMethodName         = "/testing.testpb.v1.TestService/PingList"
	TestService_PingStream_FullMethodName       = "/testing.testpb.v1.TestService/PingStream"
	TestService_PingClientStream_FullMethodName = "/testing.testpb.v1.TestService/PingClientStream"
)

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *PingEmptyRequest, opts ...grpc.CallOption) (*PingEmptyResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingErrorRequest, opts ...grpc.CallOption) (*PingErrorResponse, error)
	PingList(ctx context.Context, in *PingListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PingListResponse], error)
	PingStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PingStreamRequest, PingStreamResponse], error)
	PingClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PingClientStreamRequest, PingClientStreamResponse], error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	_ = "STUB: not implemented"
	return *new(TestServiceClient)
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *PingEmptyRequest, opts ...grpc.CallOption) (*PingEmptyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingErrorRequest, opts ...grpc.CallOption) (*PingErrorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PingListResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_PingListClient = grpc.ServerStreamingClient[PingListResponse]

func (c *testServiceClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PingStreamRequest, PingStreamResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_PingStreamClient = grpc.BidiStreamingClient[PingStreamRequest, PingStreamResponse]

func (c *testServiceClient) PingClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PingClientStreamRequest, PingClientStreamResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_PingClientStreamClient = grpc.ClientStreamingClient[PingClientStreamRequest, PingClientStreamResponse]

type TestServiceServer interface {
	PingEmpty(context.Context, *PingEmptyRequest) (*PingEmptyResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingErrorRequest) (*PingErrorResponse, error)
	PingList(*PingListRequest, grpc.ServerStreamingServer[PingListResponse]) error
	PingStream(grpc.BidiStreamingServer[PingStreamRequest, PingStreamResponse]) error
	PingClientStream(grpc.ClientStreamingServer[PingClientStreamRequest, PingClientStreamResponse]) error
	mustEmbedUnimplementedTestServiceServer()
}

type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) PingEmpty(context.Context, *PingEmptyRequest) (*PingEmptyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) PingError(context.Context, *PingErrorRequest) (*PingErrorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) PingList(*PingListRequest, grpc.ServerStreamingServer[PingListResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) PingStream(grpc.BidiStreamingServer[PingStreamRequest, PingStreamResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) PingClientStream(grpc.ClientStreamingServer[PingClientStreamRequest, PingClientStreamResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedTestServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_PingListServer = grpc.ServerStreamingServer[PingListResponse]

func _TestService_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_PingStreamServer = grpc.BidiStreamingServer[PingStreamRequest, PingStreamResponse]

func _TestService_PingClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_PingClientStreamServer = grpc.ClientStreamingServer[PingClientStreamRequest, PingClientStreamResponse]

var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.testpb.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingStream",
			Handler:       _TestService_PingStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PingClientStream",
			Handler:       _TestService_PingClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}
