package prometheus

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	exemplarFromCtxFn func(ctx context.Context) prometheus.Labels
	labelsFromCtxFn   func(metadata context.Context) prometheus.Labels
)

type serverMetricsConfig struct {
	counterOpts counterOptions

	histogramOpts histogramOptions

	enableHistogram bool

	contextLabels []string
}

type ServerMetricsOption func(*serverMetricsConfig)

func (c *serverMetricsConfig) apply(opts []ServerMetricsOption) { _ = "STUB: not implemented"; return }

func WithServerCounterOptions(opts ...CounterOption) ServerMetricsOption {
	_ = "STUB: not implemented"
	return *new(ServerMetricsOption)
}

func WithServerHandlingTimeHistogram(opts ...HistogramOption) ServerMetricsOption {
	_ = "STUB: not implemented"
	return *new(ServerMetricsOption)
}

func WithContextLabels(labelNames ...string) ServerMetricsOption {
	_ = "STUB: not implemented"
	return *new(ServerMetricsOption)
}
