package gorequest

import (
	"encoding/json"
	"go.dtapp.net/gostring"
	"log"
)

// Params 参数
type Params map[string]interface{}

// NewParams 新建参数
func NewParams() Params {
	P := make(Params)
	return P
}

// NewParamsWith 参数使用
func NewParamsWith(params ...Params) Params {
	p := make(Params)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

// Set 设置参数
func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

// SetParams 批量设置参数
func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

// GetParamsString 获取参数字符串
func GetParamsString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return gostring.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() map[string]interface{} {
	targetMap := make(map[string]interface{})

	// 从原始复制到目标
	for key, value := range *p {
		targetMap[key] = value
	}

	// 重新申请一个新的map
	*p = map[string]interface{}{}
	return targetMap
}
