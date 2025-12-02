package builder

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/core"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/echo/internal/internal/param"
)

type Route[T any] struct {
	params *param.Route[T]
	logger log.Logger
	setter kernel.Setter
}

func NewRoute[T any](picker kernel.Picker[T], setter kernel.Setter, logger log.Logger) *Route[T] {
	return &Route[T]{
		params: param.NewRoute(picker),
		logger: logger,
		setter: setter,
	}
}

func (r *Route[T]) Validator(validator kernel.Validator[T]) (route *Route[T]) {
	r.params.Validator = validator
	route = r

	return
}

func (r *Route[T]) Binder(binder kernel.Binder[T]) (route *Route[T]) {
	r.params.Binder = binder
	route = r

	return
}

func (r *Route[T]) Defaulter(defaulter kernel.Defaulter[T]) (route *Route[T]) {
	r.params.Defaulter = defaulter
	route = r

	return
}

func (r *Route[T]) Build() *core.Router[T] {
	return core.NewRouter(r.params, r.setter, r.logger)
}
