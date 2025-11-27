package kernel

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func NewContext(context echo.Context) *Context {
	return &Context{
		Context: context,
	}
}

func (c *Context) Fill(target any) (err error) {
	if be := c.Bind(target); nil != be {
		err = be
	} else if ve := c.Validate(target); nil != ve {
		err = ve
	}

	return
}
