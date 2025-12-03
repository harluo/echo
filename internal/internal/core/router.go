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
	setter  kernel.Setter
	handler *internal.Handler[T]
}

func NewRouter[T any](
	handler kernel.Handler[T, any], params *param.Route[T],
	setter kernel.Setter, logger log.Logger,
) *Router[T] {
	return &Router[T]{
		params:  params,
		setter:  setter,
		handler: internal.NewHandler[T](handler, params, logger),
	}
}

func (r *Router[T]) Get() *echo.Route {
	return r.setter.GET(r.params.Path, r.handler.Handle(), r.params.Middles...)
}

func (r *Router[T]) Put() *echo.Route {
	return r.setter.PUT(r.params.Path, r.handler.Handle(), r.params.Middles...)
}

func (r *Router[T]) Post() *echo.Route {
	return r.setter.POST(r.params.Path, r.handler.Handle(), r.params.Middles...)
}

func (r *Router[T]) Delete() *echo.Route {
	return r.setter.DELETE(r.params.Path, r.handler.Handle(), r.params.Middles...)
}

func (r *Router[T]) Options() *echo.Route {
	return r.setter.OPTIONS(r.params.Path, r.handler.Handle(), r.params.Middles...)
}
