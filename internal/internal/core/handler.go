package core

import (
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/mengpo"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	logger log.Logger
}

func newHandler(logger log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h Handler) Handle(creator kernel.CreatorFunc, handler kernel.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		request := creator() // 每次创建请求
		fields := gox.Fields[any]{
			field.New("request", request),
		}
		h.logger.Debug("收到请求", fields[0], fields[1:]...)

		if be := ctx.Bind(request); nil != be {
			err = be
			errors := fields.Add(field.Error(be))
			h.logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if bhe := (&echo.DefaultBinder{}).BindHeaders(ctx, request); nil != bhe {
			err = bhe
			errors := fields.Add(field.Error(err))
			h.logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if me := mengpo.New().Build().Set(request); nil != me {
			err = me
			errors := fields.Add(field.Error(me))
			h.logger.Warn("设置默认值出错", errors[0], errors[1:]...)
		} else if ve := ctx.Validate(request); nil != ve {
			err = ve
			errors := fields.Add(field.Error(ve))
			h.logger.Warn("数据验证出错", errors[0], errors[1:]...)
		} else if rsp, he := handler(kernel.NewContext(ctx, h.logger), request); nil != he {
			err = he
		} else {
			err = ctx.JSON(http.StatusOK, rsp)
		}

		return
	}
}
