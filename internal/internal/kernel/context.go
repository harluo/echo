package kernel

import (
	"github.com/goexl/log"
	"github.com/labstack/echo/v4"
)

type Context struct {
	context echo.Context
	logger  log.Logger
}

func NewContext(
	context echo.Context,
	logger log.Logger,
) *Context {
	return &Context{
		context: context,
		logger:  logger,
	}
}
