package internal

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newValidator,
		newLogger,
	).Build().Apply()
}
