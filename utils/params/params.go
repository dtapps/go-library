package params

import (
	"encoding/json"
	"github.com/nilorg/sdk/convert"
	"net/url"
)

type Params map[string]interface{}

func NewParams() Params {
	P := make(Params)
	return P
}

func NewParamsWithType(params ...Params) Params {
	p := make(Params)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

func (p Params) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, GetParamsString(v))
	}
	return u.Encode()
}

func GetParamsString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}
