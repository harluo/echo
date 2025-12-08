package core

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/echo/internal/core/internal"
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
	for _, binder := range router.Routes(s) {
		binder.bind()
	}
	for _, optional := range routers {
		for _, binder := range optional.Routes(s) {
			binder.bind()
		}
	}

	return s.echo.StartServer(s.server.Http())
}

func (s *Server) Stop(ctx context.Context) (err error) {
	if see := s.echo.Shutdown(ctx); nil != see {
		err = see
	} else if she := s.server.Http().Shutdown(ctx); nil != she {
		err = she
	}

	return
}

func (s *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *Group {
	return NewGroup(s.echo.Group(prefix, middles...), s.logger)
}

func (s *Server) get() kernel.Setter {
	return s.echo.GET
}

func (s *Server) post() kernel.Setter {
	return s.echo.POST
}

func (s *Server) put() kernel.Setter {
	return s.echo.PUT
}

func (s *Server) delete() kernel.Setter {
	return s.echo.DELETE
}

func (s *Server) options() kernel.Setter {
	return s.echo.OPTIONS
}

func (s *Server) getLogger() log.Logger {
	return s.logger
}
