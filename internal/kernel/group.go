package kernel

import (
	"github.com/labstack/echo/v4"
)

type Group struct {
	group *echo.Group
}

func NewGroup(group *echo.Group) *Group {
	return &Group{
		group: group,
	}
}

func (g *Group) Get(path string, handler HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return g.group.GET(path, func(ctx echo.Context) error {
		return handler(NewContext(ctx))
	}, middles...)
}

func (g *Group) Put(path string, handler HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return g.group.PUT(path, func(ctx echo.Context) error {
		return handler(NewContext(ctx))
	}, middles...)
}

func (g *Group) Post(path string, handler HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return g.group.POST(path, func(ctx echo.Context) error {
		return handler(NewContext(ctx))
	}, middles...)
}

func (g *Group) Delete(path string, handler HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return g.group.DELETE(path, func(ctx echo.Context) error {
		return handler(NewContext(ctx))
	}, middles...)
}

func (g *Group) Options(path string, handler HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return g.group.OPTIONS(path, func(ctx echo.Context) error {
		return handler(NewContext(ctx))
	}, middles...)
}
