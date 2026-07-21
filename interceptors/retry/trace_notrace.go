//go:build retrynotrace

package retry

import (
	"context"
	"fmt"
)

type notrace struct{}

func (notrace) LazyLog(x fmt.Stringer, sensitive bool) { _ = "STUB: not implemented"; return }
func (notrace) LazyPrintf(format string, a ...any)     { _ = "STUB: not implemented"; return }
func (notrace) SetError()                              { _ = "STUB: not implemented"; return }
func (notrace) SetRecycler(f func(any))                { _ = "STUB: not implemented"; return }
func (notrace) SetTraceInfo(traceID, spanID uint64)    { _ = "STUB: not implemented"; return }
func (notrace) SetMaxEvents(m int)                     { _ = "STUB: not implemented"; return }
func (notrace) Finish()                                { _ = "STUB: not implemented"; return }

func traceFromCtx(ctx context.Context) (notrace, bool) {
	_ = "STUB: not implemented"
	return *new(notrace), false
}
