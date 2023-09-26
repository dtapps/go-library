package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PowerCardInfoResponse struct {
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

type PowerCardInfoResult struct {
	Result PowerCardInfoResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newPowerCardInfoResult(result PowerCardInfoResponse, body []byte, http gorequest.Response) *PowerCardInfoResult {
	return &PowerCardInfoResult{Result: result, Body: body, Http: http}
}

// PowerCardInfo 电费充值卡详情
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (c *Client) PowerCardInfo(ctx context.Context, notMustParams ...gorequest.Params) (*PowerCardInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/cardInfo", params)
	if err != nil {
		return newPowerCardInfoResult(PowerCardInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PowerCardInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPowerCardInfoResult(response, request.ResponseBody, request), err
}
