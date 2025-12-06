package kernel

const (
	MethodGet Method = iota + 1
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
	MethodHead
	MethodOptions
	MethodTrace
)

type Method uint8
