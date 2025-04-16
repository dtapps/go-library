package framework

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

// HertzValidator Hertz框架验证器
type HertzValidator struct {
	once        sync.Once           // 用于确保初始化只执行一次
	validate    *validator.Validate // 验证器实例
	validateTag string              // 验证标签的名称，默认是 "binding"
}

// NewHertzValidator 创建并返回一个新的 Validator 实例
func NewHertzValidator() *HertzValidator {
	vd := &HertzValidator{}
	vd.lazyinit() // 初始化验证器
	return vd
}

// SliceValidationError 是一个切片类型，用于存储多个错误
type SliceValidationError []error

// Error 将所有错误信息拼接成一个字符串，用换行符分隔
func (err SliceValidationError) Error() string {
	n := len(err)
	switch n {
	case 0:
		return ""
	default:
		var b strings.Builder
		if err[0] != nil {
			fmt.Fprintf(&b, "[%d]: %s", 0, err[0].Error())
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				if err[i] != nil {
					b.WriteString("\n")
					fmt.Fprintf(&b, "[%d]: %s", i, err[i].Error())
				}
			}
		}
		return b.String()
	}
}

// ValidateStruct 验证传入的对象是否符合规则
func (m *HertzValidator) ValidateStruct(obj interface{}) error {
	if obj == nil {
		return nil
	}

	// 获取对象的反射值
	value := reflect.Value{}
	if val, ok := obj.(reflect.Value); ok {
		value = val
		obj = val.Interface()
	} else {
		value = reflect.ValueOf(obj)
	}

	// 在验证之前处理默认值
	if err := m.setDefaults(value); err != nil {
		return err
	}

	// 根据对象的类型进行不同的处理
	switch value.Kind() {
	case reflect.Ptr:
		// 如果是指针类型，递归验证指针指向的值
		return m.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		// 如果是结构体类型，调用 validateStruct 进行验证
		return m.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		// 如果是切片或数组类型，逐个验证元素
		count := value.Len()
		validateRet := make(SliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := m.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		// 其他类型（如基本类型）不进行验证
		return nil
	}
}

// setDefaults 设置字段的默认值
func (m *HertzValidator) setDefaults(value reflect.Value) error {
	// 如果是指针类型，获取指针指向的值
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// 如果不是结构体类型，直接返回
	if value.Kind() != reflect.Struct {
		return nil
	}

	// 获取结构体的类型信息
	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i) // 获取字段的值
		fieldType := t.Field(i) // 获取字段的类型信息

		// 获取字段的标签信息
		tag := fieldType.Tag.Get(m.ValidateTag())
		if tag == "" {
			continue
		}

		// 解析标签中的默认值
		defaultValue := parseDefaultValue(tag)
		if defaultValue == "" {
			continue
		}

		// 如果字段为空且可以设置值，则设置默认值
		if field.CanSet() && isEmptyValue(field) {
			if err := setFieldValue(field, defaultValue); err != nil {
				return fmt.Errorf("为字段 '%s' 设置默认值失败: %w", fieldType.Name, err)
			}
		}
	}
	return nil
}

// validateStruct 验证结构体类型的对象
func (m *HertzValidator) validateStruct(obj interface{}) error {
	m.lazyinit()                  // 确保验证器已初始化
	err := m.validate.Struct(obj) // 使用验证器进行验证
	if err == nil {
		return nil
	}

	// 处理验证错误
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return CustomValidationError(ve)
	}
	return err
}

// Engine 返回底层的验证器实例
func (m *HertzValidator) Engine() interface{} {
	m.lazyinit()
	return m.validate
}

// lazyinit 初始化验证器
func (m *HertzValidator) lazyinit() {
	if m.validate == nil {
		m.once.Do(func() {
			m.validate = validator.New()     // 创建新的验证器实例
			m.validate.SetTagName("binding") // 设置验证标签名称
			m.validateTag = "binding"
		})
	}
}

// ValidateTag 返回当前的验证标签名称
func (m *HertzValidator) ValidateTag() string {
	return m.validateTag
}

// SetValidateTag 设置验证标签名称
func (m *HertzValidator) SetValidateTag(tag string) {
	vdEngine := m.Engine().(*validator.Validate)
	vdEngine.SetTagName(tag) // 更新验证器的标签名称
	m.validateTag = tag      // 更新本地的标签名称
}
