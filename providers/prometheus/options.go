package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func FromError(err error) *status.Status { _ = "STUB: not implemented"; return nil }

type CounterOption func(*prometheus.CounterOpts)

type counterOptions []CounterOption

func (co counterOptions) apply(o prometheus.CounterOpts) prometheus.CounterOpts {
	_ = "STUB: not implemented"
	return *new(prometheus.CounterOpts)
}

func WithConstLabels(labels prometheus.Labels) CounterOption {
	_ = "STUB: not implemented"
	return *new(CounterOption)
}

func WithSubsystem(subsystem string) CounterOption {
	_ = "STUB: not implemented"
	return *new(CounterOption)
}

func WithNamespace(namespace string) CounterOption {
	_ = "STUB: not implemented"
	return *new(CounterOption)
}

type HistogramOption func(*prometheus.HistogramOpts)

type histogramOptions []HistogramOption

func (ho histogramOptions) apply(o *prometheus.HistogramOpts) prometheus.HistogramOpts {
	_ = "STUB: not implemented"
	return *new(prometheus.HistogramOpts)
}

func WithHistogramBuckets(buckets []float64) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

func WithHistogramOpts(opts *prometheus.HistogramOpts) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

func WithHistogramConstLabels(labels prometheus.Labels) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

func WithHistogramSubsystem(subsystem string) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

func WithHistogramNamespace(namespace string) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

func typeFromMethodInfo(mInfo *grpc.MethodInfo) grpcType {
	_ = "STUB: not implemented"
	return *new(grpcType)
}

type Option func(*config)

type config struct {
	exemplarFn exemplarFromCtxFn
	labelsFn   labelsFromCtxFn
}

func (c *config) apply(opts []Option) { _ = "STUB: not implemented"; return }

func WithExemplarFromContext(exemplarFn exemplarFromCtxFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLabelsFromContext(labelsFn labelsFromCtxFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
