package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaSecurityApplySetOrderPathInfoResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
}

type WxaSecurityApplySetOrderPathInfoResult struct {
	Result WxaSecurityApplySetOrderPathInfoResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newWxaSecurityApplySetOrderPathInfoResult(result WxaSecurityApplySetOrderPathInfoResponse, body []byte, http gorequest.Response) *WxaSecurityApplySetOrderPathInfoResult {
	return &WxaSecurityApplySetOrderPathInfoResult{Result: result, Body: body, Http: http}
}

// WxaSecurityApplySetOrderPathInfo 申请设置订单页 path 信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/basic-info-management/applySetOrderPathInfo.html
func (c *Client) WxaSecurityApplySetOrderPathInfo(ctx context.Context, notMustParams ...gorequest.Params) (*WxaSecurityApplySetOrderPathInfoResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/security/applysetorderpathinfo?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaSecurityApplySetOrderPathInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaSecurityApplySetOrderPathInfoResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaSecurityApplySetOrderPathInfoResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 61042:
		return "批量提交超过最大数量，一次提交的 appid 数量不超过100个"
	case 61043:
		return "参数填写错误"
	case 61044:
		return "path填写不规范"
	}
	return "系统繁忙"
}
