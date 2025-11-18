package core

import (
	"github.com/harluo/echo/internal/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func newServer(config *config.Server) (server *Server, err error) {
	e := echo.New()
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

func (server *Server) Get(path string, handler echo.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.GET(path, handler, middles...)
}

func (server *Server) Put(path string, handler echo.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.PUT(path, handler, middles...)
}

func (server *Server) Post(path string, handler echo.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.POST(path, handler, middles...)
}

func (server *Server) Delete(path string, handler echo.HandlerFunc, middles ...echo.MiddlewareFunc) *echo.Route {
	return server.echo.DELETE(path, handler, middles...)
}
