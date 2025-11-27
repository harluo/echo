package core

import (
	"net/http"

	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func newHandler() *Handler {
	return new(Handler)
}

func (h Handler) Handle(handler kernel.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if rsp, he := handler(kernel.NewContext(ctx)); nil != he {
			err = he
		} else {
			err = ctx.JSON(http.StatusOK, rsp)
		}

		return
	}
}
