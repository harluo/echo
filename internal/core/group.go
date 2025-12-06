package core

import (
	"github.com/goexl/log"
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

func (g *Group) get() kernel.Setter {
	return g.group.GET
}

func (g *Group) post() kernel.Setter {
	return g.group.POST
}

func (g *Group) put() kernel.Setter {
	return g.group.PUT
}

func (g *Group) delete() kernel.Setter {
	return g.group.DELETE
}

func (g *Group) options() kernel.Setter {
	return g.group.OPTIONS
}
