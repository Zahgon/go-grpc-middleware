package protovalidate

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type options struct {
	ignoreMessages []protoreflect.FullName
}

type Option func(*options)

func evaluateOpts(opts []Option) *options { _ = "STUB: not implemented"; return nil }

func WithIgnoreMessages(msgs ...protoreflect.MessageType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func (o *options) shouldIgnoreMessage(fqn protoreflect.FullName) bool {
	_ = "STUB: not implemented"
	return false
}
