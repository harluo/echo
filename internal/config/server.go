package config

import (
	"fmt"
)

type Server struct {
	// 接口
	Interface string `json:"interface,omitempty"`
	// 端口
	Port uint16 `default:"9090" json:"port,omitempty" validate:"min=1000,max=65535"`
	// 超时
	Timeout *Timeout `json:"timeout,omitempty"`
}

func newServer(http *Http) *Server {
	return http.Server
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Interface, s.Port)
}
