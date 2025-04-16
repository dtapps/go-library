package framework

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func (c *Context) ginBindPathParams(ctx *gin.Context, obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("obj 必须是非 nil 的指针")
	}
	v = v.Elem()
	t := v.Type()

	hasPath := false // 记录是否有 path 标签
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if pathTag := field.Tag.Get("path"); pathTag != "" {
			hasPath = true
			break
		}
	}

	// 如果结构体没有 path 字段，则直接返回
	if !hasPath {
		return nil
	}

	// 有 path 字段，进行路径参数绑定
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		pathTag := field.Tag.Get("path")
		if pathTag == "" {
			continue
		}

		paramVal := ctx.Param(pathTag)
		if paramVal == "" {
			if def := field.Tag.Get("default"); def != "" {
				paramVal = def
			} else {
				return fmt.Errorf("路径参数 '%s' 不能为空", pathTag)
			}
		}

		fv := v.Field(i)
		if !fv.CanSet() {
			return fmt.Errorf("字段 %s 不可设置", field.Name)
		}

		switch fv.Kind() {
		case reflect.String:
			fv.SetString(paramVal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(paramVal, 10, 64)
			if err != nil {
				return fmt.Errorf("字段 %s 类型转换失败: %w", field.Name, err)
			}
			fv.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(paramVal, 10, 64)
			if err != nil {
				return fmt.Errorf("字段 %s 类型转换失败: %w", field.Name, err)
			}
			fv.SetUint(uv)
		case reflect.Bool:
			bv, err := strconv.ParseBool(paramVal)
			if err != nil {
				return fmt.Errorf("字段 %s 类型转换失败: %w", field.Name, err)
			}
			fv.SetBool(bv)
		default:
			return fmt.Errorf("字段 %s 类型 %s 不支持", field.Name, fv.Kind())
		}
	}
	return nil
}
