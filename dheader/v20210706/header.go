package v20210706

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Headers map[string]interface{}

func NewHeaders() Headers {
	P := make(Headers)
	return P
}

func (p Headers) Set(key string, value interface{}) {
	p[key] = value
}

func (p Headers) SetHeaders(headers Headers) {
	for key, value := range headers {
		p[key] = value
	}
}

func (p Headers) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, GetHeadersString(v))
	}
	return u.Encode()
}

func GetHeadersString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
