package govalidator

import (
	"errors"
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

type ValidatorConfig struct {
	Local        string
	DefaultPage  string
	DefaultLimit string
}

type Validator struct {
	config   ValidatorConfig
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidator(config *ValidatorConfig) (*Validator, error) {

	var err error
	v := &Validator{config: *config}

	// 获取gin的校验器
	var ok bool
	if v.validate, ok = binding.Validator.Engine().(*validator.Validate); ok {

		// 注册自定义标签
		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 参数名称
			paramsNameTag := field.Tag.Get("params_name")
			if paramsNameTag == "-" {
				// 将大写的User替换为json中定义的tag标签 -- "LoginForm.user": "user长度不能超过10个字符"
				oldName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				if oldName == "-" {
					return ""
				}
			}
			return paramsNameTag
		})

		// 注册自定义标签
		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 参数名称
			pageTag := field.Tag.Get("page")
			if pageTag == "-" {
				// 将大写的User替换为json中定义的tag标签 -- "LoginForm.user": "user长度不能超过10个字符"
				oldName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				if oldName == "-" {
					return v.config.DefaultPage
				}
			}
			return pageTag
		})

		// 注册自定义标签
		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 参数名称
			limitTag := field.Tag.Get("limit")
			if limitTag == "-" {
				// 将大写的User替换为json中定义的tag标签 -- "LoginForm.user": "user长度不能超过10个字符"
				oldName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				if oldName == "-" {
					return v.config.DefaultLimit
				}
			}
			return limitTag
		})

		zhT := zh.New() // 中文
		enT := en.New() // 英文
		v.uni = ut.New(enT, zhT, enT)

		var o bool
		v.trans, o = v.uni.GetTranslator(v.config.Local)
		if !o {
			return nil, nil
		}

		// 注册翻译器
		switch v.config.Local {
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
	var errs validator.ValidationErrors
	errors.As(err, &errs)
	for _, err := range errs {
		errMsg = err.Translate(v.trans)
		break
	}
	return
}
