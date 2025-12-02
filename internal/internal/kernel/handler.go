package kernel

type Handler[Q any, P any] func(*Context, Q) (P, error)
