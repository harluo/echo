package kernel

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	context echo.Context
}

func NewContext(context echo.Context) *Context {
	return &Context{
		context: context,
	}
}

func (c *Context) Fill(target any) (err error) {
	if be := c.context.Bind(target); nil != be {
		err = be
	} else if ve := c.context.Validate(target); nil != ve {
		err = ve
	}

	return
}
