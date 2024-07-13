package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOilCardInfoResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Id       string `json:"id"`        // 充值卡ID，用于电费推单
		CardNum  string `json:"card_num"`  // 用户电费户号
		Name     string `json:"name"`      // 姓名
		Phone    string `json:"phone"`     // 手机号
		UserNum  string `json:"user_num"`  // 身份证
		CardType string `json:"card_type"` // 0中石化 1中石油
	} `json:"data"`
}

type RestOilCardInfoResult struct {
	Result RestOilCardInfoResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newRestOilCardInfoResult(result RestOilCardInfoResponse, body []byte, http gorequest.Response) *RestOilCardInfoResult {
	return &RestOilCardInfoResult{Result: result, Body: body, Http: http}
}

// RestOilCardInfo 油卡充值卡详情
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/373
func (c *Client) RestOilCardInfo(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilCardInfoResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Oil/cardInfo")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response RestOilCardInfoResponse
	request, err := c.request(ctx, "rest/Oil/cardInfo", params, &response)
	return newRestOilCardInfoResult(response, request.ResponseBody, request), err
}
