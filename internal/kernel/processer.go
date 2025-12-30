package kernel

type Processer interface {
	Process(*Context, Next) error
}
