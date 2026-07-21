package timeout

import (
	"time"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}
