package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerDelCardResponse struct {
	Data string `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

type RestPowerDelCardResult struct {
	Result RestPowerDelCardResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestPowerDelCardResult(result RestPowerDelCardResponse, body []byte, http gorequest.Response) *RestPowerDelCardResult {
	return &RestPowerDelCardResult{Result: result, Body: body, Http: http}
}

// RestPowerDelCard 删除电费充值卡
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/330
func (c *Client) RestPowerDelCard(ctx context.Context, cardID int64, notMustParams ...gorequest.Params) (*RestPowerDelCardResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Power/delCard")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("card_id", cardID)

	// 请求
	var response RestPowerDelCardResponse
	request, err := c.request(ctx, "rest/Power/delCard", params, &response)
	return newRestPowerDelCardResult(response, request.ResponseBody, request), err
}
