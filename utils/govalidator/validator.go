package govalidator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type Validator struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidator(local string) (*Validator, error) {

	v := &Validator{}

	// 获取gin的校验器
	var ok bool
	if v.validate, ok = binding.Validator.Engine().(*validator.Validate); ok {

		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 参数名称
			paramsName := field.Tag.Get("params_name")
			if paramsName == "-" {
				// 将大写的User替换为json中定义的tag标签 -- "LoginForm.user": "user长度不能超过10个字符"
				oldName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				if oldName == "-" {
					return ""
				}
			}
			return paramsName
		})

		zhT := zh.New() // 中文
		enT := en.New() // 英文
		v.uni = ut.New(enT, zhT, enT)

		var o bool
		v.trans, o = v.uni.GetTranslator(local)
		if !o {
			return nil, nil
		}

		// 注册翻译器
		var err error
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v.validate, v.trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v.validate, v.trans)
		default:
			err = chTranslations.RegisterDefaultTranslations(v.validate, v.trans)
		}

		if err != nil {
			return nil, err
		}

		return v, nil
	}
	return nil, nil
}

// Translate 翻译错误信息
// https://learnku.com/articles/59498
// https://www.cnblogs.com/silent-cxl/p/15181647.html
func (v *Validator) Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(v.trans)
		break
	}
	return
}
