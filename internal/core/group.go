package core

import (
	"github.com/goexl/log"
	"github.com/goexl/validate"
	"github.com/harluo/echo/internal/internal/util"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Group struct {
	group     *echo.Group
	validator validate.Validator
	logger    log.Logger
}

func NewGroup(group *echo.Group, validator validate.Validator, logger log.Logger) *Group {
	return &Group{
		group:     group,
		validator: validator,
		logger:    logger,
	}
}

func (g *Group) Group(prefix string, processers ...kernel.Processer) *Group {
	middles := make([]echo.MiddlewareFunc, 0, len(processers))
	for _, processer := range processers {
		middles = append(middles, util.NewProcesser(processer))
	}

	return &Group{
		group: g.group.Group(prefix, middles...),
	}
}

func (g *Group) get() util.Setter {
	return g.group.GET
}

func (g *Group) post() util.Setter {
	return g.group.POST
}

func (g *Group) put() util.Setter {
	return g.group.PUT
}

func (g *Group) delete() util.Setter {
	return g.group.DELETE
}

func (g *Group) options() util.Setter {
	return g.group.OPTIONS
}

func (g *Group) getLogger() log.Logger {
	return g.logger
}

func (g *Group) getValidator() validate.Validator {
	return g.validator
}
