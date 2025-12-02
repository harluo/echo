package kernel

import (
	"github.com/labstack/echo/v4"
)

type Setter interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
