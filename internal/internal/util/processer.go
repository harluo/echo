package util

import (
	"github.com/harluo/echo/internal/internal/util/internal"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

func NewProcesser(processer kernel.Processer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return processer.Process(kernel.NewContext(ctx), internal.NewNext(next))
		}
	}
}
