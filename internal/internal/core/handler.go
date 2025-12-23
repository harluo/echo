package core

import (
	"encoding/json"
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
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

func (h *Handler[Q, S]) Handle(logger log.Logger) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		context := kernel.NewContext(ctx)

		request := new(Q) // 每次创建请求
		fields := gox.Fields[any]{
			field.New("url", ctx.Request().URL),
			field.New("method", ctx.Request().Method),
			field.New("request", request),
		}
		logger.Debug("收到请求", fields[0], fields[1:]...)

		if be := h.params.Binder(context, request); nil != be {
			errors := fields.Add(field.Error(be))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    1,
				"message": "数据格式有错误",
			})
			logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if me := h.params.Defaulter(context, request); nil != me {
			errors := fields.Add(field.Error(me))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    2,
				"message": "数据不匹配",
			})
			logger.Warn("设置默认值出错", errors[0], errors[1:]...)
		} else if ve := h.params.Validator(context, request); nil != ve {
			errors := fields.Add(field.Error(ve))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    3,
				"message": "数据检查无效",
			})
			logger.Warn("数据验证出错", errors[0], errors[1:]...)
		} else if rsp, he := h.handler(context, request); nil != he {
			err = h.handleError(ctx, he)
		} else if nil == rsp {
			err = h.handleException(ctx, request)
		} else {
			err = h.handleSuccess(ctx, rsp)
		}

		return
	}
}

func (h *Handler[Q, S]) handleSuccess(ctx echo.Context, response *S) error {
	code := http.StatusOK

	method := ctx.Request().Method
	switch method {
	case "POST":
		code = gox.Ift(h.params.Asynchronous, http.StatusAccepted, http.StatusCreated)
	}

	return ctx.JSON(code, response)
}

func (h *Handler[Q, S]) handleException(ctx echo.Context, request *Q) error {
	code := http.StatusNotFound

	method := ctx.Request().Method
	switch method {
	case "GET":
		code = http.StatusNoContent
	default:
		code = http.StatusNoContent
	}

	return ctx.JSON(code, map[string]any{
		"message": "无返回数据",
		"data":    request,
	})
}

func (h *Handler[Q, S]) handleError(ctx echo.Context, original error) (err error) {
	switch converted := original.(type) {
	case json.Marshaler:
		if bytes, mje := converted.MarshalJSON(); nil == mje {
			err = ctx.JSONBlob(http.StatusBadGateway, bytes)
		} else {
			err = original
		}
	default:
		err = original
	}

	return
}
