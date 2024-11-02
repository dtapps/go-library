package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerAddCardResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		CardNum    string `json:"card_num"`    // 用户电费户号
		StoreId    string `json:"store_id"`    // 店铺ID
		CreateTime int    `json:"create_time"` // 创建时间
		Type       int    `json:"type"`        // 0国家电网 1南方电网
		CmsUid     int    `json:"cms_uid"`
		Province   string `json:"province"` // 省份，带省。
		City       string `json:"city"`     // 城市，带市
		Id         string `json:"id"`       // 充值卡ID，用于电费推单
		Remark     string `json:"remark"`   // 备注
	} `json:"data"`
}

type RestPowerAddCardResult struct {
	Result RestPowerAddCardResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestPowerAddCardResult(result RestPowerAddCardResponse, body []byte, http gorequest.Response) *RestPowerAddCardResult {
	return &RestPowerAddCardResult{Result: result, Body: body, Http: http}
}

// RestPowerAddCard 添加电费充值卡
// card_num = 用户电费户号
// province = 省份，带省
// city = 城市，带市
// type = 0国家电网 1南方电网
// remark = 备注
// user_ext = 南网必填，请输入用户信息，身份证后六位 / 营业执照后六位 / 银行卡后六位 ，三者选任意一个即可
// https://open.wikeyun.cn/#/apiDocument/9/document/326
func (c *Client) RestPowerAddCard(ctx context.Context, cardNum string, province string, city string, Type int64, notMustParams ...gorequest.Params) (*RestPowerAddCardResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	params.Set("card_num", cardNum)        // 用户电费户号
	params.Set("province", province)       // 省份，带省
	params.Set("city", city)               // 城市，带市
	params.Set("type", Type)               // 0国家电网 1南方电网

	// 请求
	var response RestPowerAddCardResponse
	request, err := c.request(ctx, "rest/Power/addCard", params, &response)
	return newRestPowerAddCardResult(response, request.ResponseBody, request), err
}
