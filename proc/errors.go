package proc

import "errors"

var (
	ErrAlreadyTerminated = errors.New("nolan.proc: process already terminated")
)
