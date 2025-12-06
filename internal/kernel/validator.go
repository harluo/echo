package kernel

type Validator[T any] func(*Context, *T) error
