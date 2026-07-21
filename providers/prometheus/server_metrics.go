package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerMetrics struct {
	serverStartedCounter    *prometheus.CounterVec
	serverHandledCounter    *prometheus.CounterVec
	serverStreamMsgReceived *prometheus.CounterVec
	serverStreamMsgSent     *prometheus.CounterVec

	serverHandledHistogram *prometheus.HistogramVec

	contextLabelNames []string
}

func NewServerMetrics(opts ...ServerMetricsOption) *ServerMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *ServerMetrics) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (m *ServerMetrics) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (m *ServerMetrics) InitializeMetrics(server reflection.ServiceInfoProvider) {
	_ = "STUB: not implemented"
	return
}

func (m *ServerMetrics) preRegisterMethod(serviceName string, mInfo *grpc.MethodInfo) {
	_ = "STUB: not implemented"
	return
}

func (m *ServerMetrics) UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (m *ServerMetrics) StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
