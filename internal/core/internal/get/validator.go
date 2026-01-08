package get

import (
	"github.com/goexl/validate"
	"github.com/harluo/di"
)

type Validator struct {
	di.Get

	Validator validate.Validator `optional:"true"`
}
