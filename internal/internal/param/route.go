package param

import (
	"github.com/go-playground/validator/v10"
	"github.com/goexl/mengpo"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Route[T any] struct {
	Picker    kernel.Picker[T]
	Validator kernel.Validator[T]
	Binder    kernel.Binder[T]
	Defaulter kernel.Defaulter[T]

	Path    string
	Middles []echo.MiddlewareFunc
}

func NewRoute[T any](picker kernel.Picker[T]) *Route[T] {
	return &Route[T]{
		Picker: picker,
		Validator: func(ctx *kernel.Context, t T) error {
			return validator.New().StructCtx(ctx, t)
		},
		Binder: func(ctx *kernel.Context, t T) (err error) {
			if be := ctx.Bind(t); nil != be {
				err = be
			} else if bhe := (&echo.DefaultBinder{}).BindHeaders(ctx.Echo(), t); nil != bhe {
				err = bhe
			}

			return
		},
		Defaulter: func(ctx *kernel.Context, t T) error {
			return mengpo.New().Build().Set(t)
		},
		Path:    "",
		Middles: make([]echo.MiddlewareFunc, 0),
	}
}
