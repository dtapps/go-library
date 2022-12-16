package wechatopen

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
		CountryCode     string `json:"countryCode"`     // 区号
		Watermark       struct {
			Timestamp int    `json:"timestamp"` // 用户获取手机号操作的时间戳
			Appid     string `json:"appid"`     // 小程序appid
		} `json:"watermark"` // 数据水印
	} `json:"phone_info"` // 用户手机号信息
}

type WxaBusinessGetUserPhoneNumberResult struct {
	Result WxaBusinessGetUserPhoneNumberResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
}

func newWxaBusinessGetUserPhoneNumberResult(result WxaBusinessGetUserPhoneNumberResponse, body []byte, http gorequest.Response) *WxaBusinessGetUserPhoneNumberResult {
	return &WxaBusinessGetUserPhoneNumberResult{Result: result, Body: body, Http: http}
}

// WxaBusinessGetUserPhoneNumber code换取用户手机号。 每个 code 只能使用一次，code的有效期为5min
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (c *Client) WxaBusinessGetUserPhoneNumber(ctx context.Context, code string) (*WxaBusinessGetUserPhoneNumberResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	params.Set("code", code)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/business/getuserphonenumber?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaBusinessGetUserPhoneNumberResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaBusinessGetUserPhoneNumberResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaBusinessGetUserPhoneNumberResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85013:
		return "无效的自定义配置"
	case 85014:
		return "无效的模板编号"
	case 85043:
		return "模板错误"
	case 85044:
		return "代码包超过大小限制"
	case 85045:
		return "ext_json 有不存在的路径"
	case 85046:
		return "tabBar 中缺少 path"
	case 85047:
		return "pages 字段为空"
	case 85048:
		return "ext_json 解析失败"
	case 80082:
		return "没有权限使用该插件"
	case 80067:
		return "找不到使用的插件"
	case 80066:
		return "非法的插件版本"
	case 9402202:
		return "请勿频繁提交，待上一次操作完成后再提交"
	case 9402203:
		return `标准模板ext_json错误，传了不合法的参数， 如果是标准模板库的模板，则ext_json支持的参数仅为{"extAppid":'', "ext": {}, "window": {}}`
	}
	return "系统繁忙"
}
