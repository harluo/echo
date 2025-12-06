package kernel

import (
	"github.com/labstack/echo/v4"
)

type Setter func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
