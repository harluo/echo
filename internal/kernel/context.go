package kernel

import (
	"context"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

var _ context.Context = (*Context)(nil)

type Context struct {
	echo echo.Context
	ctx  context.Context
}

func NewContext(echo echo.Context) *Context {
	return &Context{
		echo: echo,
		ctx:  context.Background(),
	}
}

func (c *Context) Echo() echo.Context {
	return c.echo
}

func (c *Context) Deadline() (time.Time, bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *Context) Header(key string) string {
	return c.echo.Request().Header.Get(key)
}

func (c *Context) Body() (body *[]byte, err error) {
	if bytes, rae := io.ReadAll(c.echo.Request().Body); rae != nil {
		err = rae
	} else {
		body = &bytes
	}

	return
}

func (c *Context) Method() string {
	return c.echo.Request().Method
}

func (c *Context) Bind(target any) error {
	return c.echo.Bind(target)
}

func (c *Context) With(key string, value any) (ctx *Context) {
	c.ctx = context.WithValue(c.ctx, key, value)
	ctx = c

	return
}
