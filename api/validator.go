package api

import (
	"relinc/util"

	"github.com/go-playground/validator/v10"
)

var validCurrrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurency((currency))
	}
	return false
}
