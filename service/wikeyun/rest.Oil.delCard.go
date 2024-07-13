package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOilDelCardResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

type RestOilDelCardResult struct {
	Result RestOilDelCardResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newRestOilDelCardResult(result RestOilDelCardResponse, body []byte, http gorequest.Response) *RestOilDelCardResult {
	return &RestOilDelCardResult{Result: result, Body: body, Http: http}
}

// RestOilDelCard 删除油卡充值卡
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/372
func (c *Client) RestOilDelCard(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilDelCardResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Oil/delCard")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response RestOilDelCardResponse
	request, err := c.request(ctx, "rest/Oil/delCard", params, &response)
	return newRestOilDelCardResult(response, request.ResponseBody, request), err
}
