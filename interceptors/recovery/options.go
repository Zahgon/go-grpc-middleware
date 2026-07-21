package recovery

var defaultOptions = &options{
	recoveryHandlerFunc: nil,
}

type options struct {
	recoveryHandlerFunc RecoveryHandlerFuncContext
}

func evaluateOptions(opts []Option) *options { _ = "STUB: not implemented"; return nil }

type Option func(*options)

func WithRecoveryHandler(f RecoveryHandlerFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRecoveryHandlerContext(f RecoveryHandlerFuncContext) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
