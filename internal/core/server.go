package core

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/validate"
	"github.com/goexl/validator"
	"github.com/harluo/di"
	"github.com/harluo/echo/internal/core/internal"
	"github.com/harluo/echo/internal/core/internal/get"
	"github.com/harluo/echo/internal/internal/util"
	"github.com/harluo/httpd"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo      *echo.Echo
	http      *httpd.Server
	validator validate.Validator
	logger    log.Logger
}

func newServer(
	http *httpd.Server,
	logger *internal.Logger,
) (server *Server) {
	server = new(Server)
	server.http = http
	server.logger = logger.Logger()

	e := echo.New()
	e.HideBanner = true                      // 禁用标志输出
	e.Logger = logger                        // 日志
	e.HTTPErrorHandler = server.errorHandler // 日志
	server.echo = e
	if ie := di.New().Instance().Get(server.detectValidator).Build().Inject(); ie != nil { // 注入校验器
		server.validator = validator.New()
	}

	return
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

	return s.echo.StartServer(s.http.Http())
}

func (s *Server) Stop(ctx context.Context) (err error) {
	if see := s.echo.Shutdown(ctx); nil != see {
		err = see
	} else if she := s.http.Http().Shutdown(ctx); nil != she {
		err = she
	}

	return
}

func (s *Server) Group(prefix string, middles ...echo.MiddlewareFunc) *Group {
	return NewGroup(s.echo.Group(prefix, middles...), s.validator, s.logger)
}

func (s *Server) errorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	content := []byte(err.Error())
	switch converted := err.(type) {
	case json.Marshaler:
		if bytes, mje := converted.MarshalJSON(); nil == mje {
			content = bytes
		} else {
			content = []byte(mje.Error())
		}
	case *echo.HTTPError:
		code = converted.Code
	}

	if je := c.JSONBlob(code, content); je != nil {
		s.logger.Error("错误处理出错", field.Error(je))
	}
}

func (s *Server) get() util.Setter {
	return s.echo.GET
}

func (s *Server) post() util.Setter {
	return s.echo.POST
}

func (s *Server) put() util.Setter {
	return s.echo.PUT
}

func (s *Server) delete() util.Setter {
	return s.echo.DELETE
}

func (s *Server) options() util.Setter {
	return s.echo.OPTIONS
}

func (s *Server) getLogger() log.Logger {
	return s.logger
}

func (s *Server) getValidator() validate.Validator {
	return s.validator
}

func (s *Server) detectValidator(gv get.Validator) {
	if detected := gv.Validator; detected == nil {
		s.validator = validator.New()
	}
}
