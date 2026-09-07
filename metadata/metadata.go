package metadata

import (
	"context"

	grpcMetadata "google.golang.org/grpc/metadata"
)

type MD grpcMetadata.MD

func ExtractIncoming(ctx context.Context) MD { _ = "STUB: not implemented"; return *new(MD) }

func ExtractOutgoing(ctx context.Context) MD { _ = "STUB: not implemented"; return *new(MD) }

func (m MD) Clone(copiedKeys ...string) MD { _ = "STUB: not implemented"; return *new(MD) }

func (m MD) ToOutgoing(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (m MD) ToIncoming(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (m MD) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (m MD) Del(key string) MD { _ = "STUB: not implemented"; return *new(MD) }

func (m MD) Set(key, value string) MD { _ = "STUB: not implemented"; return *new(MD) }

func (m MD) Add(key, value string) MD { _ = "STUB: not implemented"; return *new(MD) }
