package testvalidatev1

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	TestValidateService_Send_FullMethodName       = "/testing.testvalidate.v1.TestValidateService/Send"
	TestValidateService_SendStream_FullMethodName = "/testing.testvalidate.v1.TestValidateService/SendStream"
)

type TestValidateServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	SendStream(ctx context.Context, in *SendStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SendStreamResponse], error)
}

type testValidateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestValidateServiceClient(cc grpc.ClientConnInterface) TestValidateServiceClient {
	_ = "STUB: not implemented"
	return *new(TestValidateServiceClient)
}

func (c *testValidateServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testValidateServiceClient) SendStream(ctx context.Context, in *SendStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SendStreamResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestValidateService_SendStreamClient = grpc.ServerStreamingClient[SendStreamResponse]

type TestValidateServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	SendStream(*SendStreamRequest, grpc.ServerStreamingServer[SendStreamResponse]) error
}

type UnimplementedTestValidateServiceServer struct{}

func (UnimplementedTestValidateServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestValidateServiceServer) SendStream(*SendStreamRequest, grpc.ServerStreamingServer[SendStreamResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestValidateServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeTestValidateServiceServer interface {
	mustEmbedUnimplementedTestValidateServiceServer()
}

func RegisterTestValidateServiceServer(s grpc.ServiceRegistrar, srv TestValidateServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _TestValidateService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestValidateService_SendStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestValidateService_SendStreamServer = grpc.ServerStreamingServer[SendStreamResponse]

var TestValidateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.testvalidate.v1.TestValidateService",
	HandlerType: (*TestValidateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _TestValidateService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendStream",
			Handler:       _TestValidateService_SendStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "testing/testvalidate/v1/test_validate.proto",
}
