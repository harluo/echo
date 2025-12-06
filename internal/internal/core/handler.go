package core

import (
	"encoding/json"
	"net/http"

	"github.com/harluo/echo/internal/internal/param"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Handler[Q any, S any] struct {
	handler kernel.Handler[Q, S]
	params  *param.Route[Q]
}

func NewHandler[Q any, S any](handler kernel.Handler[Q, S], params *param.Route[Q]) *Handler[Q, S] {
	return &Handler[Q, S]{
		handler: handler,
		params:  params,
	}
}

func (h *Handler[Q, S]) Handle() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		context := kernel.NewContext(ctx)

		request := new(Q) // 每次创建请求
		/* todo fields := gox.Fields[any]{
			field.New("request", request),
		}*/
		// todo h.logger.Debug("收到请求", fields[0], fields[1:]...)

		if be := h.params.Binder(context, request); nil != be {
			// todo errors := fields.Add(field.Error(be))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    1,
				"message": "数据格式有错误",
			})
			// todo h.logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if me := h.params.Defaulter(context, request); nil != me {
			// todo errors := fields.Add(field.Error(me))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    2,
				"message": "数据不匹配",
			})
			// todo h.logger.Warn("设置默认值出错", errors[0], errors[1:]...)
		} else if ve := h.params.Validator(context, request); nil != ve {
			// todo errors := fields.Add(field.Error(ve))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    3,
				"message": "数据无效",
			})
			// todo h.logger.Warn("数据验证出错", errors[0], errors[1:]...)
		} else if rsp, he := h.handler(context, request); nil != he {
			err = h.handleException(ctx, he)
		} else {
			err = ctx.JSON(http.StatusOK, rsp)
		}

		return
	}
}

func (h *Handler[Q, S]) handleException(ctx echo.Context, exception error) (err error) {
	switch converted := exception.(type) {
	case json.Marshaler:
		if bytes, mje := converted.MarshalJSON(); nil == mje {
			err = ctx.JSONBlob(http.StatusBadGateway, bytes)
		} else {
			err = exception
		}
	default:
		err = exception
	}

	return
}
