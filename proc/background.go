package proc

import (
	"context"
)

// BootableProcess a kind of process which runs when Start is called and lasts until Stop routine is called.
type BootableProcess interface {
	Start() error
	Stop(ctx context.Context) error
}
