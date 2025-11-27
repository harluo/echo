package core

import (
	"github.com/harluo/echo/internal/internal/core"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Group struct {
	group   *echo.Group
	handler *core.Handler
}

func NewGroup(group *echo.Group, handler *core.Handler) *Group {
	return &Group{
		group:   group,
		handler: handler,
	}
}

func (g *Group) Get(
	path string,
	creator kernel.CreatorFunc, handler kernel.HandlerFunc,
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return g.group.GET(path, g.handler.Handle(creator, handler), middles...)
}

func (g *Group) Put(
	path string,
	creator kernel.CreatorFunc, handler kernel.HandlerFunc,
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return g.group.PUT(path, g.handler.Handle(creator, handler), middles...)
}

func (g *Group) Post(
	path string,
	creator kernel.CreatorFunc, handler kernel.HandlerFunc,
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return g.group.POST(path, g.handler.Handle(creator, handler), middles...)
}

func (g *Group) Delete(
	path string,
	creator kernel.CreatorFunc, handler kernel.HandlerFunc,
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return g.group.DELETE(path, g.handler.Handle(creator, handler), middles...)
}

func (g *Group) Options(
	path string,
	creator kernel.CreatorFunc, handler kernel.HandlerFunc,
	middles ...echo.MiddlewareFunc,
) *echo.Route {
	return g.group.OPTIONS(path, g.handler.Handle(creator, handler), middles...)
}
