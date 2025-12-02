package core

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/builder"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Group struct {
	group  *echo.Group
	logger log.Logger
}

func NewGroup(group *echo.Group, logger log.Logger) *Group {
	return &Group{
		group:  group,
		logger: logger,
	}
}

func (g *Group) Group(prefix string, middles ...echo.MiddlewareFunc) *Group {
	return &Group{
		group: g.group.Group(prefix, middles...),
	}
}

func (g *Group) Route(picker kernel.Picker[any]) *builder.Route[any] {
	return builder.NewRoute(picker, g.group, g.logger)
}
