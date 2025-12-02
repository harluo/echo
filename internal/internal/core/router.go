package core

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/core/internal"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/labstack/echo/v4"
)

type Router[T any] struct {
	setter  kernel.Setter
	handler *internal.Handler[T]
}

func NewRouter[T any](params *param.Route[T], setter kernel.Setter, logger log.Logger) *Router[T] {
	return &Router[T]{
		setter:  setter,
		handler: internal.NewHandler[T](params, logger),
	}
}

func (r *Router[T]) Get(
	path string,
	handler kernel.Handler[T, any],
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return r.setter.GET(path, r.handler.Handle(handler), middles...)
}

func (r *Router[T]) Put(
	path string,
	handler kernel.Handler[T, any],
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return r.setter.PUT(path, r.handler.Handle(handler), middles...)
}

func (r *Router[T]) Post(
	path string,
	handler kernel.Handler[T, any],
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return r.setter.POST(path, r.handler.Handle(handler), middles...)
}

func (r *Router[T]) Delete(
	path string,
	handler kernel.Handler[T, any],
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return r.setter.DELETE(path, r.handler.Handle(handler), middles...)
}

func (r *Router[T]) Options(
	path string,
	handler kernel.Handler[T, any],
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return r.setter.OPTIONS(path, r.handler.Handle(handler), middles...)
}
