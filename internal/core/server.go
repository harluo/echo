package core

import (
	"context"

	"github.com/harluo/echo/internal/internal/core"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/httpd"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo    *echo.Echo
	server  *httpd.Server
	handler *core.Handler
}

func newServer(server *httpd.Server, handler *core.Handler) *Server {
	e := echo.New()
	e.HideBanner = true // 禁用标志输出

	return &Server{
		echo:    e,
		server:  server,
		handler: handler,
	}
}

func (s *Server) Start(_ context.Context, router Router, routers ...Router) error {
	router.Route(s)
	for _, optional := range routers {
		optional.Route(s)
	}

	return s.echo.StartServer(s.server.Server)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *Group {
	return NewGroup(s.echo.Group(prefix, middles...), s.handler)
}

func (s *Server) Get(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.GET(path, s.handler.Handle(handler), middles...)
}

func (s *Server) Put(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.PUT(path, s.handler.Handle(handler), middles...)
}

func (s *Server) Post(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.POST(path, s.handler.Handle(handler), middles...)
}

func (s *Server) Delete(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.DELETE(path, s.handler.Handle(handler), middles...)
}

func (s *Server) Options(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return s.echo.OPTIONS(path, s.handler.Handle(handler), middles...)
}
