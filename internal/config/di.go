package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newHttp,

		newServer,
		newTimeout,
	).Build().Apply()
}
