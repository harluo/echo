package kernel

type Binder[T any] func(*Context, *T) error
