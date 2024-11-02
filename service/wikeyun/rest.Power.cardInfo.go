package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerCardInfoResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id       string `json:"id"`       // 充值卡ID，用于电费推单
		CardNum  string `json:"card_num"` // 用户电费户号
		Province string `json:"province"` // 省份，带省。
		City     string `json:"city"`     // 城市，带市
		StoreId  string `json:"store_id"` // 店铺ID
		Type     int    `json:"type"`     // 0国家电网 1南方电网
		Remark   string `json:"remark"`
	} `json:"data"`
}

type RestPowerCardInfoResult struct {
	Result RestPowerCardInfoResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newRestPowerCardInfoResult(result RestPowerCardInfoResponse, body []byte, http gorequest.Response) *RestPowerCardInfoResult {
	return &RestPowerCardInfoResult{Result: result, Body: body, Http: http}
}

// RestPowerCardInfo 电费充值卡详情
// card_id = 充值卡ID
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (c *Client) RestPowerCardInfo(ctx context.Context, cardID int64, notMustParams ...gorequest.Params) (*RestPowerCardInfoResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("card_id", cardID) // 充值卡ID

	// 请求
	var response RestPowerCardInfoResponse
	request, err := c.request(ctx, "rest/Power/cardInfo", params, &response)
	return newRestPowerCardInfoResult(response, request.ResponseBody, request), err
}
