package prometheus

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/prometheus/client_golang/prometheus"
)

type reporter struct {
	clientMetrics   *ClientMetrics
	serverMetrics   *ServerMetrics
	typ             interceptors.GRPCType
	service, method string
	kind            Kind
	exemplar        prometheus.Labels
	contextLabels   []string
}

func (r *reporter) PostCall(err error, rpcDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) PostMsgSend(_ any, _ error, sendDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) PostMsgReceive(_ any, _ error, recvDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

type reportable struct {
	clientMetrics *ClientMetrics
	serverMetrics *ServerMetrics

	opts []Option
}

func (rep *reportable) ServerReporter(ctx context.Context, meta interceptors.CallMeta) (interceptors.Reporter, context.Context) {
	_ = "STUB: not implemented"
	return *new(interceptors.Reporter), *new(context.Context)
}

func (rep *reportable) ClientReporter(ctx context.Context, meta interceptors.CallMeta) (interceptors.Reporter, context.Context) {
	_ = "STUB: not implemented"
	return *new(interceptors.Reporter), *new(context.Context)
}

func (rep *reportable) reporter(ctx context.Context, sm *ServerMetrics, cm *ClientMetrics, meta interceptors.CallMeta, kind Kind) (interceptors.Reporter, context.Context) {
	_ = "STUB: not implemented"
	return *new(interceptors.Reporter), *new(context.Context)
}

func (r *reporter) incrementWithExemplar(c *prometheus.CounterVec, lvals ...string) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) observeWithExemplar(h *prometheus.HistogramVec, value float64, lvals ...string) {
	_ = "STUB: not implemented"
	return
}
