package auth

import (
	"context"
)

const (
	headerAuthorize = "authorization"
)

func AuthFromMD(ctx context.Context, expectedScheme string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
