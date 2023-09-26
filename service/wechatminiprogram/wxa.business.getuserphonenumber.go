package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaBusinessGetUserPhoneNumberResponse struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
		PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
		CountryCode     int    `json:"countryCode"`     // 区号
		Watermark       struct {
			Timestamp int    `json:"timestamp"` // 用户获取手机号操作的时间戳
			Appid     string `json:"appid"`     // 小程序appid
		} `json:"watermark"`
	} `json:"phone_info"`
}

type WxaBusinessGetUserPhoneNumberResult struct {
	Result WxaBusinessGetUserPhoneNumberResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
}

func newWxaBusinessGetUserPhoneNumberResult(result WxaBusinessGetUserPhoneNumberResponse, body []byte, http gorequest.Response) *WxaBusinessGetUserPhoneNumberResult {
	return &WxaBusinessGetUserPhoneNumberResult{Result: result, Body: body, Http: http}
}

// WxaBusinessGetUserPhoneNumber code换取用户手机号。 每个code只能使用一次，code的有效期为5min
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (c *Client) WxaBusinessGetUserPhoneNumber(ctx context.Context, notMustParams ...gorequest.Params) (*WxaBusinessGetUserPhoneNumberResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/business/getuserphonenumber?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newWxaBusinessGetUserPhoneNumberResult(WxaBusinessGetUserPhoneNumberResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaBusinessGetUserPhoneNumberResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaBusinessGetUserPhoneNumberResult(response, request.ResponseBody, request), err
}
