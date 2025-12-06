package core

import (
	"github.com/harluo/echo/internal/internal/core"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type RouteDefault[Q any, S any] struct {
	target  Target
	params  *param.Route[Q]
	handler *core.Handler[Q, S]
}

func NewRouteDefault[Q any, S any](
	target Target, handler kernel.Handler[Q, S],
	params *param.Route[Q],
) *RouteDefault[Q, S] {
	return &RouteDefault[Q, S]{
		target:  target,
		params:  params,
		handler: core.NewHandler[Q, S](handler, params),
	}
}

func (r *RouteDefault[Q, S]) bind() {
	var route *echo.Route
	switch r.params.Method {
	case kernel.MethodGet:
		route = r.target.get()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	case kernel.MethodPost:
		route = r.target.post()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	case kernel.MethodPut:
		route = r.target.put()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	case kernel.MethodDelete:
		route = r.target.delete()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	case kernel.MethodOptions:
		route = r.target.options()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	default:
		route = r.target.get()(r.params.Path, r.handler.Handle(), r.params.Middles...)
	}
	route.Name = r.params.Name

	return
}
