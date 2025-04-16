package framework

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"sync"
)

// Validator 验证数据
func (c *Context) Validator(obj any) error {
	validate := getValidator() // 复用验证器实例
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
