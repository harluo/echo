package core

import (
	"github.com/goexl/xiren"
)

type Validator struct {
	// 无字段
}

func newValidator() *Validator {
	return new(Validator)
}

func (v *Validator) Validate(target any) (err error) {
	if ve := xiren.Struct(target); nil != ve {
		err = ve
	}

	return
}
