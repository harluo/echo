package kernel

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/mengpo"
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

func (c *Context) Fill(req any, purpose string) (err error) {
	fields := gox.Fields[any]{
		field.New("purpose", purpose),
		field.New("request", req),
	}
	c.logger.Debug("收到请求", fields[0], fields[1:]...)

	if be := c.context.Bind(req); nil != be {
		err = be
		errors := fields.Add(field.Error(be))
		c.logger.Warn("绑定值出错", errors[0], errors[1:]...)
	} else if bhe := (&echo.DefaultBinder{}).BindHeaders(c.context, req); nil != bhe {
		err = bhe
		errors := fields.Add(field.Error(err))
		c.logger.Warn("绑定值出错", errors[0], errors[1:]...)
	} else if me := mengpo.New().Build().Set(req); nil != me {
		err = me
		errors := fields.Add(field.Error(me))
		c.logger.Warn("设置默认值出错", errors[0], errors[1:]...)
	} else if ve := c.context.Validate(req); nil != ve {
		err = ve
		errors := fields.Add(field.Error(ve))
		c.logger.Warn("数据验证出错", errors[0], errors[1:]...)
	}

	return
}
