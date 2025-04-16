package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"reflect"
	"strconv"
)

func (c *Context) ginBindPathParams(ginCtx *gin.Context, obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("obj 必须是非 nil 的指针")
	}
	v = v.Elem()
	t := v.Type()

	hasPath := false
	fields := make([]reflect.StructField, 0)

	// 第一次循环：检查是否有 path 标签的字段，记录它们
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if pathTag := field.Tag.Get("path"); pathTag != "" {
			hasPath = true
			fields = append(fields, field)
		}
	}

	// 如果没有 path 字段，直接返回
	if !hasPath {
		return nil
	}

	// 第二次循环：进行路径参数的绑定
	for _, field := range fields {
		pathTag := field.Tag.Get("path")
		paramVal := ginCtx.Param(pathTag)
		if paramVal == "" {
			if def := field.Tag.Get("default"); def != "" {
				paramVal = def
			} else {
				return fmt.Errorf("路径参数 '%s' 不能为空", pathTag)
			}
		}

		fv := v.FieldByName(field.Name)
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

func (c *Context) ginBindJson(ginCtx *gin.Context, obj any) error {
	bodyBytes, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return fmt.Errorf("读取请求体失败: %w", err)
	}
	defer ginCtx.Request.Body.Close()

	// 空 body 直接返回，不报错
	if len(bodyBytes) == 0 {
		return nil
	}

	log.Println("原始 Body:", string(bodyBytes))

	// 反序列化
	if err := json.Unmarshal(bodyBytes, obj); err != nil {
		return fmt.Errorf("JSON 反序列化失败: %w", err)
	}

	// 如果后续还需要用 Body，重新设置回来
	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return nil
}
