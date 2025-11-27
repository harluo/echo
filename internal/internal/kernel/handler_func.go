package kernel

type HandlerFunc func(*Context) (any, error)
