package core

import (
	"github.com/goexl/log"
	"github.com/harluo/echo/internal/internal/kernel"
)

type Target interface {
	get() kernel.Setter

	post() kernel.Setter

	put() kernel.Setter

	delete() kernel.Setter

	options() kernel.Setter

	getLogger() log.Logger
}
