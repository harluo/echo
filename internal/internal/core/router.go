package core

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/core/internal"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/labstack/echo/v4"
)

type Router[T any] struct {
	params  *param.Route[T]
	setter  kernel.Targeter
	handler *internal.Handler[T]
}

func NewRouter[T any](
	handler kernel.Handler[T, any], params *param.Route[T],
	setter kernel.Targeter, logger log.Logger,
) *Router[T] {
	return &Router[T]{
		params:  params,
		setter:  setter,
		handler: internal.NewHandler[T](handler, params, logger),
	}
}

func (r *Router[T]) Get() (route *echo.Route) {
	route = r.setter.GET(r.params.Path, r.handler.Handle(), r.params.Middles...)
	route.Name = r.params.Name

	return
}

func (r *Router[T]) Put() (route *echo.Route) {
	route = r.setter.PUT(r.params.Path, r.handler.Handle(), r.params.Middles...)
	route.Name = r.params.Name

	return
}

func (r *Router[T]) Post() (route *echo.Route) {
	route = r.setter.POST(r.params.Path, r.handler.Handle(), r.params.Middles...)
	route.Name = r.params.Name

	return
}

func (r *Router[T]) Delete() (route *echo.Route) {
	route = r.setter.DELETE(r.params.Path, r.handler.Handle(), r.params.Middles...)
	route.Name = r.params.Name

	return
}

func (r *Router[T]) Options() (route *echo.Route) {
	route = r.setter.OPTIONS(r.params.Path, r.handler.Handle(), r.params.Middles...)
	route.Name = r.params.Name

	return
}
