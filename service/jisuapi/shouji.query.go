package jisuapi

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
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

type ShoujiQueryResult struct {
	Result ShoujiQueryResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newShoujiQueryResult(result ShoujiQueryResponse, body []byte, http gorequest.Response) *ShoujiQueryResult {
	return &ShoujiQueryResult{Result: result, Body: body, Http: http}
}

// ShoujiQuery 手机号码归属地
// https://www.jisuapi.com/api/shouji/
func (c *Client) ShoujiQuery(ctx context.Context, shouji string, appkey string, notMustParams ...gorequest.Params) (*ShoujiQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "shouji/query")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("shouji", shouji) // 手机号

	// 请求
	var response ShoujiQueryResponse
	request, err := c.request(ctx, "shouji/query?appkey="+appkey, params, http.MethodGet, &response)
	return newShoujiQueryResult(response, request.ResponseBody, request), err
}
