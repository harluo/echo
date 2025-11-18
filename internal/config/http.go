package config

import (
	"github.com/harluo/config"
)

type Http struct {
	Server *Server `default:"{}" json:"server,omitempty" validate:"required"`
}

func newHttp(config config.Getter) (http *Http, err error) {
	http = new(Http)
	err = config.Get(&struct {
		Http *Http `default:"{}" json:"http,omitempty" validate:"required"`
	}{
		Http: http,
	})

	return
}
