package framework

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// Param 获取路径参数
func (c *Context) Param(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Param(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Param(key)
	}
	return ""
}

func (c *Context) Query(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Query(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Query(key)
	}
	return ""
}

func (c *Context) DefaultQuery(key, defaultValue string) string {
	if c.ginCtx != nil {
		return c.ginCtx.DefaultQuery(key, defaultValue)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.DefaultQuery(key, defaultValue)
	}
	return ""
}
func (c *Context) GetQuery(key string) (string, bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetQuery(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetQuery(key)
	}
	return "", false
}

//func (c *Context) QueryArray(key string) (values []string) {
//	if c.ginCtx != nil {
//		return c.ginCtx.QueryArray(key)
//	}
//	if c.hertzCtx != nil {
//		return c.hertzCtx.QueryArgs(key)
//	}
//	return
//}

func (c *Context) PostForm(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.PostForm(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.PostForm(key)
	}
	return ""
}

func (c *Context) DefaultPostForm(key, defaultValue string) string {
	if c.ginCtx != nil {
		return c.ginCtx.DefaultPostForm(key, defaultValue)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.DefaultPostForm(key, defaultValue)
	}
	return ""
}

func (c *Context) PostFormArray(key string) (values []string) {
	if c.ginCtx != nil {
		return c.ginCtx.PostFormArray(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.PostFormArray(key)
	}
	return
}

func (c *Context) GetPostForm(key string) (string, bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetPostForm(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetPostForm(key)
	}
	return "", false
}

func (c *Context) GetPostFormArray(key string) (values []string, ok bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetPostFormArray(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetPostFormArray(key)
	}
	return
}

// BindAndValidate 统一绑定参数并校验
func (c *Context) BindAndValidate(obj any) error {
	var bindErr error

	if c.ginCtx != nil {
		// Gin 会自动 bind JSON、Query、Form、Path
		bindErr = c.ginCtx.ShouldBind(obj)
	}
	if c.hertzCtx != nil {
		// Hertz 根据 Content-Type 来自动选择绑定的方法，其中 GET 请求会调用 BindQuery, 带有 Body 的请求会根据 Content-Type 自动选择
		if err := c.hertzCtx.BindByContentType(obj); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	}

	if bindErr != nil {
		return fmt.Errorf("参数绑定失败: %w", bindErr)
	}

	// 设置默认值并验证
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
	// 设置默认值
	if err := setDefaultValues(obj); err != nil {
		return err
	}

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
		errMap["general"] = fmt.Sprintf("An unexpected error occurred: %s", err.Error())
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

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 检查是否有 default 标签
		defaultValue := fieldType.Tag.Get("default")
		if defaultValue == "" {
			continue
		}

		// 跳过非可导出字段
		if !field.CanSet() {
			continue
		}

		// 如果字段已被赋值，跳过默认值设置
		if !field.IsZero() {
			continue
		}

		// 根据字段类型设置默认值
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := strconv.ParseInt(defaultValue, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid default value for field '%s': %w", fieldType.Name, err)
			}
			field.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(defaultValue, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid default value for field '%s': %w", fieldType.Name, err)
			}
			field.SetUint(val)
		case reflect.Float32, reflect.Float64:
			val, err := strconv.ParseFloat(defaultValue, 64)
			if err != nil {
				return fmt.Errorf("invalid default value for field '%s': %w", fieldType.Name, err)
			}
			field.SetFloat(val)
		case reflect.String:
			field.SetString(defaultValue)
		case reflect.Bool:
			val, err := strconv.ParseBool(defaultValue)
			if err != nil {
				return fmt.Errorf("invalid default value for field '%s': %w", fieldType.Name, err)
			}
			field.SetBool(val)
		case reflect.Slice:
			// 如果 default 值是一个切片字符串，例如 "[1,2,3]"，需要解析为具体的切片类型
			if strings.HasPrefix(defaultValue, "[") && strings.HasSuffix(defaultValue, "]") {
				// 移除前后的方括号
				cleaned := defaultValue[1 : len(defaultValue)-1]
				// 根据切片的类型来进行解析（这里只做了 int 类型的解析示例）
				elemType := fieldType.Type.Elem().Kind()
				switch elemType {
				case reflect.Int:
					// 解析为 []int
					elems := strings.Split(cleaned, ",")
					var intSlice []int
					for _, elem := range elems {
						elem = strings.TrimSpace(elem)
						val, err := strconv.Atoi(elem)
						if err != nil {
							return fmt.Errorf("invalid default value for slice field '%s': %w", fieldType.Name, err)
						}
						intSlice = append(intSlice, val)
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.String:
					// 解析为 []string
					elems := strings.Split(cleaned, ",")
					var strSlice []string
					for _, elem := range elems {
						elem = strings.TrimSpace(elem)
						strSlice = append(strSlice, elem)
					}
					field.Set(reflect.ValueOf(strSlice))
				default:
					return fmt.Errorf("unsupported slice element type '%s' for default value", elemType)
				}
			}
		default:
			return fmt.Errorf("unsupported field type '%s' for default value", field.Kind())
		}
	}

	return nil
}

// parseDefaultValue 从标签中提取默认值
func parseDefaultValue(tag string) string {
	parts := strings.Split(tag, ",") // 按逗号分割标签内容
	for _, part := range parts {
		if strings.HasPrefix(part, "default=") { // 查找以 "default=" 开头的部分
			return strings.TrimPrefix(part, "default=") // 提取默认值
		}
	}
	return ""
}

// isEmptyValue 判断字段是否为空（零值）
func isEmptyValue(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return field.String() == "" // 字符串为空
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == 0 // 整数为 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() == 0 // 无符号整数为 0
	case reflect.Float32, reflect.Float64:
		return field.Float() == 0 // 浮点数为 0
	case reflect.Bool:
		return !field.Bool() // 布尔值为 false
	case reflect.Slice, reflect.Array, reflect.Map:
		return field.Len() == 0 // 切片、数组或映射为空
	case reflect.Interface, reflect.Ptr:
		return field.IsNil() // 接口或指针为空
	default:
		return false
	}
}

// setFieldValue 根据字段的类型设置默认值
func setFieldValue(field reflect.Value, defaultValue string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(defaultValue) // 设置字符串类型的默认值
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(defaultValue, 10, 64) // 将字符串转换为整数
		if err != nil {
			return err
		}
		field.SetInt(val) // 设置整数类型的默认值
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(defaultValue, 10, 64) // 将字符串转换为无符号整数
		if err != nil {
			return err
		}
		field.SetUint(val) // 设置无符号整数类型的默认值
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(defaultValue, 64) // 将字符串转换为浮点数
		if err != nil {
			return err
		}
		field.SetFloat(val) // 设置浮点数类型的默认值
	case reflect.Bool:
		val, err := strconv.ParseBool(defaultValue) // 将字符串转换为布尔值
		if err != nil {
			return err
		}
		field.SetBool(val) // 设置布尔类型的默认值
	default:
		return fmt.Errorf("不支持的字段类型: %s", field.Kind())
	}
	return nil
}

// CustomValidationError 将验证错误格式化为自定义格式
func CustomValidationError(errs validator.ValidationErrors) error {
	var errMsgs []string
	for _, e := range errs {
		field := e.Field() // 字段名
		tag := e.Tag()     // 验证规则
		param := e.Param() // 参数值
		msg := generateErrorMessage(field, tag, param)
		errMsgs = append(errMsgs, msg)
	}
	return errors.New(strings.Join(errMsgs, ""))
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
