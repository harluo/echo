package core

import (
	"github.com/goexl/log"
	"github.com/goexl/validate"
	"github.com/harluo/echo/internal/internal/util"
)

type Target interface {
	get() util.Setter

	post() util.Setter

	put() util.Setter

	patch() util.Setter

	delete() util.Setter

	options() util.Setter

	getLogger() log.Logger

	getValidator() validate.Validator
}
