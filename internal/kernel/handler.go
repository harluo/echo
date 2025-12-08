package kernel

type Handler[Q any, S any] func(*Context, *Q) (*S, error)
