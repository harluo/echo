package kernel

import (
	"context"
	"io"
	"net/http"
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

func (c *Context) Writer() http.ResponseWriter {
	return c.echo.Response().Writer
}

func (c *Context) Queries() (query *map[string]string, err error) {
	data := make(map[string]string)
	query = &data
	binder := new(echo.DefaultBinder)
	err = binder.BindQueryParams(c.echo, query)

	return
}

func (c *Context) Bodies() (body *map[string]any, err error) {
	data := make(map[string]any)
	body = &data
	binder := new(echo.DefaultBinder)
	err = binder.BindBody(c.echo, body)

	return
}

func (c *Context) Headers() (body *map[string]string, err error) {
	data := make(map[string]string)
	body = &data
	binder := new(echo.DefaultBinder)
	err = binder.BindHeaders(c.echo, &body)

	return
}

func (c *Context) Paths() (path *map[string]any, err error) {
	data := make(map[string]any)
	path = &data
	binder := new(echo.DefaultBinder)
	err = binder.BindPathParams(c.echo, &path)

	return
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

func (c *Context) Raw() (raw *[]byte, err error) {
	if bytes, rae := io.ReadAll(c.echo.Request().Body); rae != nil {
		err = rae
	} else {
		raw = &bytes
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
