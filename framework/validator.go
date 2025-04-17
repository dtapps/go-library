package framework

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.dtapp.net/library/utils/gojson"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// BindJsonAndValidate 统一绑定Json参数并校验
func (c *Context) BindJsonAndValidate(obj any) error {
	if c.IsGin() {
		// 绑定 Path 参数
		if err := c.ginBindPathParams(c.GetGinContext(), obj); err != nil {
			return fmt.Errorf("Path 参数绑定失败：%w", err)
		}

		// 绑定 Url 参数
		if c.ginCtx.Request.URL.RawQuery != "" {
			if err := gojson.BindURLQueryToStruct(gojson.ParseURLQueryWithoutError(c.ginCtx.Request.URL.RawQuery), obj); err != nil {
				return fmt.Errorf("Url 参数绑定失败：%w", err)
			}
		}

		// 绑定 Body JSON 参数
		if err := c.ginBindJson(c.ginCtx, obj); err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("Body JSON 参数绑定失败：%w", err)
		}

		// 设置默认值
		if err := setDefaultValues(obj); err != nil {
			return fmt.Errorf("设置默认值失败：%w", err)
		}
	}
	if c.IsHertz() {
		if err := c.hertzCtx.Bind(obj); err != nil {
			return fmt.Errorf("参数绑定失败：%w", err)
		}
	}

	// 验证
	return c.Validator(obj)
}

// 验证器单例模式
var (
	globalValidator *validator.Validate
	validatorOnce   sync.Once
)

func getValidator() *validator.Validate {
	validatorOnce.Do(func() {
		globalValidator = validator.New()
	})
	return globalValidator
}

// Validator 验证数据
func (c *Context) Validator(obj any) error {
	// 复用验证器实例
	validate := getValidator()
	return validate.Struct(obj)
}

// ValidatorError 解析验证错误并生成友好的错误消息
func (c *Context) ValidatorError(err error) map[string]string {
	var errMap = make(map[string]string)
	var validationErrors validator.ValidationErrors

	// 检查是否是验证错误
	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			field := fieldErr.Field() // 字段名（保持原始大小写）
			tag := fieldErr.Tag()     // 验证规则
			param := fieldErr.Param() // 参数值
			errMap[field] = generateErrorMessage(field, tag, param)
		}
	} else if err != nil {
		// 如果不是验证错误，记录原始错误
		errMap["general"] = fmt.Sprintf("An unexpected error occurred：%s", err.Error())
	}

	return errMap
}

// setDefaultValues 设置默认值
func setDefaultValues(obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("input must be a non-nil pointer")
	}

	v = v.Elem() // 获取指针指向的值
	t := v.Type()

	// 用于字段类型到设置默认值方法的映射
	fieldSetters := map[reflect.Kind]func(field reflect.Value, defaultValue string) error{
		reflect.Int:     setDefaultInt,
		reflect.Int8:    setDefaultInt,
		reflect.Int16:   setDefaultInt,
		reflect.Int32:   setDefaultInt,
		reflect.Int64:   setDefaultInt,
		reflect.Uint:    setDefaultUint,
		reflect.Uint8:   setDefaultUint,
		reflect.Uint16:  setDefaultUint,
		reflect.Uint32:  setDefaultUint,
		reflect.Uint64:  setDefaultUint,
		reflect.Float32: setDefaultFloat,
		reflect.Float64: setDefaultFloat,
		reflect.String:  setDefaultString,
		reflect.Bool:    setDefaultBool,
		reflect.Slice:   setDefaultSlice,
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 跳过不可设置的字段
		if !field.CanSet() {
			continue
		}

		// ✅ 递归嵌套结构体（非 time.Time）
		if field.Kind() == reflect.Struct && field.Type().PkgPath() != "time" {
			// 必须传指针
			if field.CanAddr() {
				if err := setDefaultValues(field.Addr().Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// 如果没有 default 标签，跳过
		defaultValue := fieldType.Tag.Get("default")
		if defaultValue == "" {
			continue
		}

		// 如果字段已被赋值，跳过默认值设置
		if !field.IsZero() {
			continue
		}

		setter, ok := fieldSetters[field.Kind()]
		if !ok {
			return fmt.Errorf("unsupported field type '%s' for default value", field.Kind())
		}

		if err := setter(field, defaultValue); err != nil {
			return err
		}
	}

	return nil
}

func setDefaultInt(field reflect.Value, defaultValue string) error {
	val, err := strconv.ParseInt(defaultValue, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid default value for field: %w", err)
	}
	field.SetInt(val)
	return nil
}

func setDefaultUint(field reflect.Value, defaultValue string) error {
	val, err := strconv.ParseUint(defaultValue, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid default value for field: %w", err)
	}
	field.SetUint(val)
	return nil
}

func setDefaultFloat(field reflect.Value, defaultValue string) error {
	val, err := strconv.ParseFloat(defaultValue, 64)
	if err != nil {
		return fmt.Errorf("invalid default value for field: %w", err)
	}
	field.SetFloat(val)
	return nil
}

func setDefaultString(field reflect.Value, defaultValue string) error {
	field.SetString(defaultValue)
	return nil
}

func setDefaultBool(field reflect.Value, defaultValue string) error {
	val, err := strconv.ParseBool(defaultValue)
	if err != nil {
		return fmt.Errorf("invalid default value for field: %w", err)
	}
	field.SetBool(val)
	return nil
}

func setDefaultSlice(field reflect.Value, defaultValue string) error {
	if strings.HasPrefix(defaultValue, "[") && strings.HasSuffix(defaultValue, "]") {
		cleaned := defaultValue[1 : len(defaultValue)-1]
		elems := strings.Split(cleaned, ",")
		var strSlice []string
		for _, elem := range elems {
			elem = strings.TrimSpace(elem)
			strSlice = append(strSlice, elem)
		}
		field.Set(reflect.ValueOf(strSlice))
	}
	return nil
}

// generateErrorMessage 生成友好的错误消息（中文版）
func generateErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("字段 '%s' 是必填项。", field)
	case "min":
		return fmt.Sprintf("字段 '%s' 的长度必须至少为 %s 个字符。", field, param)
	case "max":
		return fmt.Sprintf("字段 '%s' 的长度不能超过 %s 个字符。", field, param)
	case "gte":
		return fmt.Sprintf("字段 '%s' 的值必须大于或等于 %s。", field, param)
	case "lte":
		return fmt.Sprintf("字段 '%s' 的值必须小于或等于 %s。", field, param)
	case "email":
		return fmt.Sprintf("字段 '%s' 必须是一个有效的电子邮件地址。", field)
	case "len":
		return fmt.Sprintf("字段 '%s' 的长度必须为 %s 个字符。", field, param)
	case "eq":
		return fmt.Sprintf("字段 '%s' 的值必须等于 %s。", field, param)
	case "ne":
		return fmt.Sprintf("字段 '%s' 的值不能等于 %s。", field, param)
	case "oneof":
		return fmt.Sprintf("字段 '%s' 的值必须是以下之一：%s。", field, param)
	default:
		return fmt.Sprintf("字段 '%s' 验证失败，规则为 '%s'。", field, tag)
	}
}
