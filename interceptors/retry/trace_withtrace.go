//go:build !retrynotrace

package retry

import (
	"context"

	t "golang.org/x/net/trace"
)

func traceFromCtx(ctx context.Context) (t.Trace, bool) {
	_ = "STUB: not implemented"
	return *new(t.Trace), false
}
