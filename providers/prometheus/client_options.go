package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type clientMetricsConfig struct {
	counterOpts counterOptions

	clientHandledHistogramFn func() *prometheus.HistogramVec

	clientStreamRecvHistogramFn func() *prometheus.HistogramVec

	clientStreamSendHistogramFn func() *prometheus.HistogramVec

	contextLabels []string
}

type ClientMetricsOption func(*clientMetricsConfig)

func (c *clientMetricsConfig) apply(opts []ClientMetricsOption) { _ = "STUB: not implemented"; return }

func WithClientCounterOptions(opts ...CounterOption) ClientMetricsOption {
	_ = "STUB: not implemented"
	return *new(ClientMetricsOption)
}

func WithClientHandlingTimeHistogram(opts ...HistogramOption) ClientMetricsOption {
	_ = "STUB: not implemented"
	return *new(ClientMetricsOption)
}

func WithClientStreamRecvHistogram(opts ...HistogramOption) ClientMetricsOption {
	_ = "STUB: not implemented"
	return *new(ClientMetricsOption)
}

func WithClientStreamSendHistogram(opts ...HistogramOption) ClientMetricsOption {
	_ = "STUB: not implemented"
	return *new(ClientMetricsOption)
}

func WithClientContextLabels(labelNames ...string) ClientMetricsOption {
	_ = "STUB: not implemented"
	return *new(ClientMetricsOption)
}
