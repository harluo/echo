package kernel

type HandlerFunc func(*Context, any) (any, error)
