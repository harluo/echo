package param

import (
	"github.com/goexl/mengpo"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Route[T any] struct {
	Name         string
	Path         string
	Method       kernel.Method
	Asynchronous bool
	Middles      []echo.MiddlewareFunc

	Initialer kernel.Initialer
	Binder    kernel.Binder[T]
	Defaulter kernel.Defaulter[T]
}

func NewRoute[T any]() *Route[T] {
	return &Route[T]{
		Name:         "",
		Path:         "",
		Method:       kernel.MethodGet,
		Asynchronous: false,
		Middles:      make([]echo.MiddlewareFunc, 0),

		Binder: func(ctx *kernel.Context, t *T) (err error) {
			if be := ctx.Bind(t); nil != be {
				err = be
			} else if bhe := (&echo.DefaultBinder{}).BindHeaders(ctx.Echo(), t); nil != bhe {
				err = bhe
			}

			return
		},
		Defaulter: func(ctx *kernel.Context, t *T) error {
			return mengpo.New().Build().Set(t)
		},
	}
}
