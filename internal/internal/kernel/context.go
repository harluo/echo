package kernel

import (
	"context"
	"time"

	"github.com/goexl/log"
	"github.com/labstack/echo/v4"
)

var _ context.Context = (*Context)(nil)

type Context struct {
	ec     echo.Context
	ctx    context.Context
	logger log.Logger
}

func NewContext(
	ec echo.Context,
	logger log.Logger,
) *Context {
	return &Context{
		ec:     ec,
		ctx:    context.Background(),
		logger: logger,
	}
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
	return c.ec.Request().Header.Get(key)
}

func (c *Context) Method() string {
	return c.ec.Request().Method
}
