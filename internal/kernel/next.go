package kernel

type Next interface {
	Execute(*Context) error
}
