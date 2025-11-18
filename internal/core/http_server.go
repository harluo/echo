package core

import (
	"net/http"

	"github.com/harluo/echo/internal/config"
)

func newHttpServer(config *config.Server) (server *http.Server) {
	server = new(http.Server)
	server.Addr = config.Addr()
	if nil != config.Timeout && 0 != config.Timeout.Read {
		server.WriteTimeout = config.Timeout.Read
	}
	if nil != config.Timeout && 0 != config.Timeout.Write {
		server.WriteTimeout = config.Timeout.Write
	}

	return
}
