package core

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/echo/internal/core/internal"
	"github.com/harluo/echo/internal/internal/builder"
	"github.com/harluo/echo/internal/internal/kernel"
	"github.com/harluo/httpd"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	server *httpd.Server
	logger log.Logger
}

func newServer(
	server *httpd.Server,
	validator *internal.Validator,
	logger *internal.Logger,
) *Server {
	e := echo.New()
	e.HideBanner = true     // 禁用标志输出
	e.Validator = validator // 校验器
	e.Logger = logger       // 日志

	return &Server{
		echo:   e,
		server: server,
		logger: logger.Logger(),
	}
}

func (s *Server) Start(_ context.Context, router Router, routers ...Router) error {
	router.Route(s)
	for _, optional := range routers {
		optional.Route(s)
	}

	return s.echo.StartServer(s.server.Http())
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *Group {
	return NewGroup(s.echo.Group(prefix, middles...), s.logger)
}

func (s *Server) Route(picker kernel.Picker[any]) *builder.Route[any] {
	return builder.NewRoute(picker, s.echo, s.logger)
}
