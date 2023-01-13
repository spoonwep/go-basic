package service

import (
	"github.com/go-playground/validator/v10"
	"go-basic/constants"
	"reflect"
)

func Validate(s any, tag ...string) validator.ValidationErrorsTranslations {
	var err error
	typ := reflect.TypeOf(s).Kind()
	if typ == reflect.Struct {
		err = constants.Validate.Struct(s)
	} else if len(tag) != 0 {
		err = constants.Validate.Var(s, tag[0])
	}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		e := errs.Translate(constants.Translator)
		return e
	}
	return nil
}
