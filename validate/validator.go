package validate

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var (
	// GlobalValidator a global general-purpose validator instance which is concurrent-safe to use.
	GlobalValidator Validator = PlaygroundValidatorAdapter{}

	playgroundValidator = validator.New()
)

// Validator is a utility component used by routines to validate a wide range of values.
type Validator interface {
	// Validate validates the given value.
	Validate(v any) error
	// ValidateWithContext validates the given value and also allows passing of context.Context for contextual
	// validation information.
	ValidateWithContext(ctx context.Context, v any) error
}

// PlaygroundValidatorAdapter is the go-playground validator adapting implementation of Validator.
// Generates nolan's exceptions if a validation error was found.
type PlaygroundValidatorAdapter struct{}

var _ Validator = PlaygroundValidatorAdapter{}

func (p PlaygroundValidatorAdapter) Validate(v any) error {
	return p.ValidateWithContext(context.Background(), v)
}

func (p PlaygroundValidatorAdapter) ValidateWithContext(ctx context.Context, v any) error {
	return AdaptValidatorError(playgroundValidator.StructCtx(ctx, v))
}
