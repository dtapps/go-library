package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpPublishResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinWxOpenQrCodeJumpPublishResult struct {
	Result CgiBinWxOpenQrCodeJumpPublishResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
	Err    error                                 // 错误
}

func newCgiBinWxOpenQrCodeJumpPublishResult(result CgiBinWxOpenQrCodeJumpPublishResponse, body []byte, http gorequest.Response, err error) *CgiBinWxOpenQrCodeJumpPublishResult {
	return &CgiBinWxOpenQrCodeJumpPublishResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxOpenQrCodeJumpPublish 发布已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumppublish.html
func (c *Client) CgiBinWxOpenQrCodeJumpPublish(ctx context.Context, prefix string) *CgiBinWxOpenQrCodeJumpPublishResult {
	// 参数
	params := gorequest.NewParams()
	params["prefix"] = prefix
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/wxopen/qrcodejumppublish?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response CgiBinWxOpenQrCodeJumpPublishResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinWxOpenQrCodeJumpPublishResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *CgiBinWxOpenQrCodeJumpPublishResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 44990:
		return "接口请求太快（超过5次/秒）"
	case 85074:
		return "小程序未发布, 小程序必须先发布代码才可以发布二维码跳转规则"
	case 85075:
		return "个人类型小程序无法设置二维码规则"
	case 85095:
		return "数据异常，请删除后重新添加"
	case 886000:
		return "本月发布次数达到上线（100次）"
	}
	return "系统繁忙"
}
