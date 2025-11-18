package core

import (
	"context"
	"net/http"

	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	server *http.Server
}

func newServer(server *http.Server) *Server {
	e := echo.New()
	e.HideBanner = true // 禁用标志输出

	return &Server{
		echo:   e,
		server: server,
	}
}

func (s *Server) Start(_ context.Context, router Router, routers ...Router) error {
	router.Route(s)
	for _, optional := range routers {
		optional.Route(s)
	}

	return s.echo.StartServer(s.server)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *kernel.Group {
	return kernel.NewGroup(s.echo.Group(prefix, middles...))
}

func (s *Server) Get(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.GET(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (s *Server) Put(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.PUT(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (s *Server) Post(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.POST(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (s *Server) Delete(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.DELETE(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (s *Server) Options(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.OPTIONS(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}
