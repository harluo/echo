package internal

import (
	"encoding/json"
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/echo/internal/internal/param"
	"github.com/labstack/echo/v4"
)

type Handler[T any] struct {
	handler kernel.Handler[T, any]
	params  *param.Route[T]
	logger  log.Logger
}

func NewHandler[T any](handler kernel.Handler[T, any], params *param.Route[T], logger log.Logger) *Handler[T] {
	return &Handler[T]{
		handler: handler,
		params:  params,
		logger:  logger,
	}
}

func (h *Handler[T]) Handle() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		context := kernel.NewContext(ctx, h.logger)

		request := h.params.Picker(context) // 每次创建请求
		fields := gox.Fields[any]{
			field.New("request", request),
		}
		h.logger.Debug("收到请求", fields[0], fields[1:]...)

		if be := h.params.Binder(context, request); nil != be {
			errors := fields.Add(field.Error(be))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    1,
				"message": "数据格式有错误",
			})
			h.logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if me := h.params.Defaulter(context, request); nil != me {
			errors := fields.Add(field.Error(me))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    2,
				"message": "数据不匹配",
			})
			h.logger.Warn("设置默认值出错", errors[0], errors[1:]...)
		} else if ve := h.params.Validator(context, request); nil != ve {
			errors := fields.Add(field.Error(ve))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    3,
				"message": "数据无效",
			})
			h.logger.Warn("数据验证出错", errors[0], errors[1:]...)
		} else if rsp, he := h.handler(context, request); nil != he {
			err = h.handleException(ctx, he)
		} else {
			err = ctx.JSON(http.StatusOK, rsp)
		}

		return
	}
}

func (h *Handler[T]) handleException(ctx echo.Context, exception error) (err error) {
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
