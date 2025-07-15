package factory

type Factory[T any] struct {
	object T
}

func New[T any](object T) *Factory[T] {
	return &Factory[T]{object: object}
}
