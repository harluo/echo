package core

import (
	"net/http"

	"github.com/goexl/log"
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

func (h Handler) Handle(handler kernel.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if rsp, he := handler(kernel.NewContext(ctx, h.logger)); nil != he {
			err = he
		} else {
			err = ctx.JSON(http.StatusOK, rsp)
		}

		return
	}
}
