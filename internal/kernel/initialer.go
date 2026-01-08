package kernel

type Initialer interface {
	Init(*Context)
}
