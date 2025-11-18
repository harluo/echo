package config

import (
	"time"
)

type Timeout struct {
	// 读数据
	Read time.Duration `json:"read,omitempty"`
	// 写数据
	Write time.Duration `json:"write,omitempty"`
}

func newTimeout(server *Server) *Timeout {
	return server.Timeout
}
