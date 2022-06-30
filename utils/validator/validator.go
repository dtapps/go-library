package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidator() (*Validator, error) {

	v := &Validator{}

	// 注册翻译器
	zhN := zh.New()
	v.uni = ut.New(zhN, zhN)

	v.trans, _ = v.uni.GetTranslator("zh")

	// 获取gin的校验器
	v.validate = binding.Validator.Engine().(*validator.Validate)

	// 注册翻译器
	err := zh_translations.RegisterDefaultTranslations(v.validate, v.trans)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Translate 翻译错误信息
// https://learnku.com/articles/59498
func (v *Validator) Translate(err error) map[string][]string {

	var result = make(map[string][]string)

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(v.trans))
	}
	return result
}
