package kernel

type Defaulter[T any] func(*Context, T) error
