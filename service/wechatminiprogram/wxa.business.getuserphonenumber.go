package wechatminiprogram

import (
	"context"
	"encoding/json"
	"fmt"
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
	Err    error                                 // 错误
}

func newWxaBusinessGetUserPhoneNumberResult(result WxaBusinessGetUserPhoneNumberResponse, body []byte, http gorequest.Response, err error) *WxaBusinessGetUserPhoneNumberResult {
	return &WxaBusinessGetUserPhoneNumberResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaBusinessGetUserPhoneNumber code换取用户手机号。 每个code只能使用一次，code的有效期为5min
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (c *Client) WxaBusinessGetUserPhoneNumber(ctx context.Context, code string) *WxaBusinessGetUserPhoneNumberResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("code", code)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/business/getuserphonenumber?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response WxaBusinessGetUserPhoneNumberResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaBusinessGetUserPhoneNumberResult(response, request.ResponseBody, request, err)
}
