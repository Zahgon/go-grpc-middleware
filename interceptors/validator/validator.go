package validator

import (
	"context"
)

type validateAller interface {
	ValidateAll() error
}

type validator interface {
	Validate(all bool) error
}

type validatorLegacy interface {
	Validate() error
}

func validate(ctx context.Context, reqOrRes interface{}, shouldFailFast bool, onValidationErrCallback OnValidationErrCallback) (err error) {
	_ = "STUB: not implemented"
	return nil
}
