package exception

type Exception interface {
	error
	Unwrap() error
}
