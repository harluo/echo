package kernel

type Next interface {
	Next(*Context) error
}
