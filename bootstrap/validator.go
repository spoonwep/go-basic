package bootstrap

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"go-basic/constants"
)

var (
	uni *ut.UniversalTranslator
)

func InitValidator() {
	translator := zh.New()
	uni = ut.New(translator, translator)
	trans, _ := uni.GetTranslator("zh")
	constants.Validate = validator.New()
	_ = zh2.RegisterDefaultTranslations(constants.Validate, trans)
	constants.Translator = trans
}
