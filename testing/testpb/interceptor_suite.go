package testpb

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	flagTls = flag.Bool("use_tls", true, "whether all gRPC middleware tests should use tls")

	certPEM []byte
	keyPEM  []byte
)

type InterceptorTestSuite struct {
	suite.Suite

	TestService TestServiceServer
	ServerOpts  []grpc.ServerOption
	ClientOpts  []grpc.DialOption

	serverAddr     string
	ServerListener net.Listener
	Server         *grpc.Server
	clientConn     *grpc.ClientConn
	Client         TestServiceClient

	restartServerWithDelayedStart chan time.Duration
	serverRunning                 chan bool

	cancels []context.CancelFunc
}

func (s *InterceptorTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

func (s *InterceptorTestSuite) RestartServer(delayedStart time.Duration) <-chan bool {
	_ = "STUB: not implemented"
	return nil
}

func (s *InterceptorTestSuite) NewClient(dialOpts ...grpc.DialOption) TestServiceClient {
	_ = "STUB: not implemented"
	return *new(TestServiceClient)
}

func (s *InterceptorTestSuite) ServerAddr() string { _ = "STUB: not implemented"; return "" }

type ctxTestNumber struct{}

var (
	ctxTestNumberKey = &ctxTestNumber{}
	zero             = 0
)

func ExtractCtxTestNumber(ctx context.Context) *int { _ = "STUB: not implemented"; return nil }

type wrappedErrFields struct {
	wrappedErr error
	fields     []any
}

func (err *wrappedErrFields) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (err *wrappedErrFields) Error() string { _ = "STUB: not implemented"; return "" }

func (err *wrappedErrFields) GRPCStatus() *status.Status { _ = "STUB: not implemented"; return nil }

func WrapFieldsInError(err error, fields []any) error { _ = "STUB: not implemented"; return nil }

func ExtractErrorFields(err error) []any { _ = "STUB: not implemented"; return nil }

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (s *InterceptorTestSuite) SimpleCtx() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *InterceptorTestSuite) DeadlineCtx(deadline time.Time) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *InterceptorTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func generateCertAndKey(san []string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
