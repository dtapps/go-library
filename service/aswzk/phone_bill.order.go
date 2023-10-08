package aswzk

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PhoneBillOrderResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   []struct {
		Id        int64  `json:"id,omitempty"`
		CityName  string `json:"city_name"`  // 地区名称
		Sort      int64  `json:"sort"`       // 排序
		Initial   string `json:"initial"`    // 首字母
		NeedYtype int64  `json:"need_ytype"` // 是否三要素认证
		NeedCity  int64  `json:"need_city"`  // 是否需要选择城市（当此开关打开以后才有下面的城市列表）
		City      []struct {
			Id       int64  `json:"id,omitempty"`
			CityName string `json:"city_name"` // 城市名称
			Initial  string `json:"initial"`   // 首字母
		} `json:"city"` // 支持的地级市
	} `json:"data,omitempty"`
}

type PhoneBillOrderResult struct {
	Result PhoneBillOrderResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newPhoneBillOrderResult(result PhoneBillOrderResponse, body []byte, http gorequest.Response) *PhoneBillOrderResult {
	return &PhoneBillOrderResult{Result: result, Body: body, Http: http}
}

// PhoneBillOrder 话费下单
func (c *Client) PhoneBillOrder(ctx context.Context, notMustParams ...gorequest.Params) (*PhoneBillOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/phone_bill/order", params, http.MethodPost)
	if err != nil {
		return newPhoneBillOrderResult(PhoneBillOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PhoneBillOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPhoneBillOrderResult(response, request.ResponseBody, request), err
}
