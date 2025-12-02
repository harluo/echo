package kernel

type Picker[T any] func(ctx *Context) T
