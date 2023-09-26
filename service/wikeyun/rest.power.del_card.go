package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
// https://open.wikeyun.cn/#/apiDocument/9/document/330
func (c *Client) RestPowerDelCard(ctx context.Context, cardId string, notMustParams ...gorequest.Params) (*RestPowerDelCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("card_id", cardId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/delCard", params)
	if err != nil {
		return newRestPowerDelCardResult(RestPowerDelCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerDelCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerDelCardResult(response, request.ResponseBody, request), err
}
