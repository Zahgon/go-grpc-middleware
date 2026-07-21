package testvalidate

import (
	"context"

	testvalidatev1 "github.com/grpc-ecosystem/go-grpc-middleware/v2/testing/testvalidate/v1"
)

type TestValidateService struct {
	testvalidatev1.UnimplementedTestValidateServiceServer
}

func (v *TestValidateService) Send(
	_ context.Context,
	_ *testvalidatev1.SendRequest,
) (*testvalidatev1.SendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *TestValidateService) SendStream(
	_ *testvalidatev1.SendStreamRequest,
	stream testvalidatev1.TestValidateService_SendStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	BadUnaryRequest = &testvalidatev1.SendRequest{
		Message: "%any",
	}

	GoodUnaryRequest = &testvalidatev1.SendRequest{
		Message: "good@example.com",
	}
)

var (
	BadStreamRequest = &testvalidatev1.SendStreamRequest{
		Message: "%any",
	}

	GoodStreamRequest = &testvalidatev1.SendStreamRequest{
		Message: "good@example.com",
	}
)
