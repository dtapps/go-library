package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOilEditCardResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

type RestOilEditCardResult struct {
	Result RestOilEditCardResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newRestOilEditCardResult(result RestOilEditCardResponse, body []byte, http gorequest.Response) *RestOilEditCardResult {
	return &RestOilEditCardResult{Result: result, Body: body, Http: http}
}

// RestOilEditCard 编辑油卡充值卡
// card_id = 充值卡ID
// card_num = 卡号
// name = 姓名
// phone = 手机号
// card_type = 类型 0中石化 1中石油
// user_num = 身份证号
// https://open.wikeyun.cn/#/apiDocument/9/document/371
func (c *Client) RestOilEditCard(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilEditCardResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response RestOilEditCardResponse
	request, err := c.request(ctx, "rest/Oil/editCard", params, &response)
	return newRestOilEditCardResult(response, request.ResponseBody, request), err
}
