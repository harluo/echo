package core

import (
	"github.com/harluo/echo/internal/config"
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func newServer(config *config.Server) (server *Server, err error) {
	e := echo.New()

	// 自定义配置
	if nil != config.Timeout && 0 != config.Timeout.Read {
		e.Server.WriteTimeout = config.Timeout.Read
	}
	if nil != config.Timeout && 0 != config.Timeout.Write {
		e.Server.WriteTimeout = config.Timeout.Write
	}
	if err = e.Start(config.Addr()); nil == err {
		server = new(Server)
		server.echo = e
	}

	return
}

func (server *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *echo.Group {
	return server.echo.Group(prefix, middles...)
}

func (server *Server) Get(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.GET(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (server *Server) Put(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.PUT(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (server *Server) Post(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.POST(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (server *Server) Delete(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.DELETE(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}

func (server *Server) Options(path string, handler kernel.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.OPTIONS(path, func(ctx echo.Context) error {
		return handler(kernel.NewContext(ctx))
	}, middles...)
}
