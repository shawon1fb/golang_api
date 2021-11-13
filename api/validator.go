package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/shawon1fb/golang_api/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
