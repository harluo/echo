package builder

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/core"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/labstack/echo/v4"
)

type Route[T any] struct {
	params  *param.Route[T]
	handler kernel.Handler[T, any]

	logger log.Logger
	setter kernel.Targeter
}

func NewRoute[T any](
	picker kernel.Picker[T], handler kernel.Handler[T, any],
	setter kernel.Targeter, logger log.Logger,
) *Route[T] {
	return &Route[T]{
		params:  param.NewRoute(picker),
		handler: handler,

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

func (r *Route[T]) Path(path string) (route *Route[T]) {
	r.params.Path = path
	route = r

	return
}

func (r *Route[T]) Name(name string) (route *Route[T]) {
	r.params.Name = name
	route = r

	return
}

func (r *Route[T]) Middleware(middleware echo.MiddlewareFunc, optionals ...echo.MiddlewareFunc) (route *Route[T]) {
	r.params.Middles = append(r.params.Middles, middleware)
	r.params.Middles = append(r.params.Middles, optionals...)
	route = r

	return
}

func (r *Route[T]) Build() *core.Router[T] {
	return core.NewRouter(r.handler, r.params, r.setter, r.logger)
}
