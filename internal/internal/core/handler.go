package core

import (
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/validate"
	"github.com/harluo/echo/internal/internal/constant"
	"github.com/harluo/echo/internal/internal/core/internal"
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

func (h *Handler[Q, S]) Handle(validator validate.Validator, logger log.Logger) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		context := kernel.NewContext(ctx)
		if initialer := h.params.Initialer; initialer != nil {
			initialer.Init(context)
		}
		if value := context.Value(internal.ContextConverter); value == nil {
			context.With("converter", h)
		}

		request := new(Q) // 每次创建请求
		fields := gox.Fields[any]{
			field.New("url", ctx.Request().URL),
			field.New("method", ctx.Request().Method),
			field.New("request", request),
		}

		if be := h.bindData(context, request); nil != be {
			errors := fields.Add(field.Error(be))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    1,
				"message": "数据格式有错误",
				"data": map[string]any{
					"error": be.Error(),
				},
			})
			logger.Warn("绑定值出错", errors[0], errors[1:]...)
		} else if me := h.setDefault(context, request); nil != me {
			errors := fields.Add(field.Error(me))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    2,
				"message": "数据不匹配",
				"data": map[string]any{
					"error": me.Error(),
				},
			})
			logger.Warn("设置默认值出错", errors[0], errors[1:]...)
		} else if ve := h.validate(context, request, validator); nil != ve {
			errors := fields.Add(field.Error(ve))
			err = ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
				"code":    3,
				"message": "数据检查无效",
				"data":    ve,
			})
			logger.Warn("数据验证出错", errors[0], errors[1:]...)
		} else if rsp, he := h.handler(context, request); nil != he {
			err = he
		} else if value := context.Value(constant.ContextResponse); value == nil && rsp != nil {
			err = h.handleSuccess(ctx, rsp)
		} else if response, ok := value.(bool); ok && response && rsp != nil {
			err = h.handleSuccess(ctx, rsp)
		} else if rsp == nil && response {
			err = h.handleException(ctx, request)
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

func (h *Handler[Q, S]) Convert(from string) string {
	return gox.String(from).Switch().Camel().Build().Case()
}

func (h *Handler[Q, S]) bindData(ctx *kernel.Context, req *Q) (err error) {
	if h.params.Binding {
		err = h.params.Binder(ctx, req)
	}

	return
}

func (h *Handler[Q, S]) setDefault(ctx *kernel.Context, req *Q) (err error) {
	if h.params.Default {
		err = h.params.Defaulter(ctx, req)
	}

	return
}

func (h *Handler[Q, S]) validate(ctx *kernel.Context, req *Q, validator validate.Validator) (err error) {
	if h.params.Validate {
		err = validator.Validate(ctx, req)
	}

	return
}
