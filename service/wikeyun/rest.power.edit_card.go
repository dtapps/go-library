package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestPowerEditCardResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type RestPowerEditCardResult struct {
	Result RestPowerEditCardResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newRestPowerEditCardResult(result RestPowerEditCardResponse, body []byte, http gorequest.Response) *RestPowerEditCardResult {
	return &RestPowerEditCardResult{Result: result, Body: body, Http: http}
}

// RestPowerEditCard 编辑电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (c *Client) RestPowerEditCard(ctx context.Context, notMustParams ...*gorequest.Params) (*RestPowerEditCardResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/editCard", params)
	if err != nil {
		return newRestPowerEditCardResult(RestPowerEditCardResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerEditCardResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerEditCardResult(response, request.ResponseBody, request), err
}
