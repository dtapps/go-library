package alipayopen

import (
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
)

func (c *Client) newParamsWithType(_method string, param ...*gorequest.Params) *gorequest.Params {
	params := gorequest.NewParamsWith(param...)
	params.Set("app_id", c.GetAppId())
	params.Set("method", _method)
	params.Set("format", "JSON")
	params.Set("charset", "utf-8")
	params.Set("sign_type", "RSA2")
	params.Set("timestamp", gotime.Current().SetFormat(gotime.DateTimeFormat))
	params.Set("version", "1.0")
	return params
}
