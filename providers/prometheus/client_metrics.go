package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

type ClientMetrics struct {
	clientStartedCounter    *prometheus.CounterVec
	clientHandledCounter    *prometheus.CounterVec
	clientStreamMsgReceived *prometheus.CounterVec
	clientStreamMsgSent     *prometheus.CounterVec

	clientHandledHistogram *prometheus.HistogramVec

	clientStreamRecvHistogram *prometheus.HistogramVec

	clientStreamSendHistogram *prometheus.HistogramVec

	contextLabelNames []string
}

func NewClientMetrics(opts ...ClientMetricsOption) *ClientMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMetrics) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (m *ClientMetrics) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (m *ClientMetrics) UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func (m *ClientMetrics) StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}
