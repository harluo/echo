package builder

import (
	"github.com/harluo/echo/internal/core"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/harluo/echo/internal/internal/util"
	"github.com/harluo/echo/internal/kernel"
)

type Route[Q any, S any] struct {
	targeter core.Target
	params   *param.Route[Q]
	handler  kernel.Handler[Q, S]
}

func NewRoute[Q any, S any](targeter core.Target, handler kernel.Handler[Q, S]) *Route[Q, S] {
	return &Route[Q, S]{
		targeter: targeter,
		params:   param.NewRoute[Q](),
		handler:  handler,
	}
}

func (r *Route[Q, S]) Get() *Route[Q, S] {
	return r.method(kernel.MethodGet)
}

func (r *Route[Q, S]) Post() *Route[Q, S] {
	return r.method(kernel.MethodPost)
}

func (r *Route[Q, S]) Put() *Route[Q, S] {
	return r.method(kernel.MethodPut)
}

func (r *Route[Q, S]) Delete() *Route[Q, S] {
	return r.method(kernel.MethodDelete)
}

func (r *Route[Q, S]) Options() *Route[Q, S] {
	return r.method(kernel.MethodOptions)
}

func (r *Route[Q, S]) Asynchronous() (route *Route[Q, S]) {
	r.params.Asynchronous = true
	route = r

	return
}

func (r *Route[Q, S]) Validator(validator kernel.Validator[Q]) (route *Route[Q, S]) {
	r.params.Validator = validator
	route = r

	return
}

func (r *Route[Q, S]) Binder(binder kernel.Binder[Q]) (route *Route[Q, S]) {
	r.params.Binder = binder
	route = r

	return
}

func (r *Route[Q, S]) Defaulter(defaulter kernel.Defaulter[Q]) (route *Route[Q, S]) {
	r.params.Defaulter = defaulter
	route = r

	return
}

func (r *Route[Q, S]) Path(path string) (route *Route[Q, S]) {
	r.params.Path = path
	route = r

	return
}

func (r *Route[Q, S]) Name(name string) (route *Route[Q, S]) {
	r.params.Name = name
	route = r

	return
}

func (r *Route[Q, S]) Middleware(required kernel.Processer, optionals ...kernel.Processer) (route *Route[Q, S]) {
	r.params.Middles = append(r.params.Middles, util.NewProcesser(required))
	for _, optional := range optionals {
		r.params.Middles = append(r.params.Middles, util.NewProcesser(optional))
	}
	route = r

	return
}

func (r *Route[Q, S]) Build() *core.RouteDefault[Q, S] {
	return core.NewRouteDefault(r.targeter, r.handler, r.params)
}

func (r *Route[Q, S]) method(method kernel.Method) (route *Route[Q, S]) {
	r.params.Method = method
	route = r

	return
}
