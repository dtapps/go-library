package jisuapi

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ShoujiQueryResponse struct {
	Status string `json:"status"` // 状态码
	Msg    string `json:"msg"`
	Result struct {
		Province string `json:"province,omitempty"` // 省
		City     string `json:"city,omitempty"`     // 市
		Company  string `json:"company,omitempty"`  // 运营商
		Cardtype string `json:"cardtype,omitempty"` // 卡类型
	} `json:"result,omitempty"`
}

// ShoujiQuery 手机号码归属地
// https://www.jisuapi.com/api/shouji/
func (c *Client) ShoujiQuery(ctx context.Context, shouji string, appkey string, notMustParams ...*gorequest.Params) (response ShoujiQueryResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("shouji", shouji) // 手机号

	// 请求
	err = c.request(ctx, "shouji/query?appkey="+appkey, params, http.MethodGet, &response)
	return
}
