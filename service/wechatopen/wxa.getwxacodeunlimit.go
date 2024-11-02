package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetUnlimitedQRCodeResponse struct {
	Errcode     int    `json:"errcode"`               // 错误码
	Errmsg      string `json:"errmsg"`                // 错误信息
	ContentType string `json:"contentType,omitempty"` // 内容类型
	Buffer      any    `json:"buffer"`                // 图片 Buffer
}

type GetUnlimitedQRCodeResult struct {
	Result GetUnlimitedQRCodeResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newGetUnlimitedQRCodeResult(result GetUnlimitedQRCodeResponse, body []byte, http gorequest.Response) *GetUnlimitedQRCodeResult {
	return &GetUnlimitedQRCodeResult{Result: result, Body: body, Http: http}
}

// GetUnlimitedQRCode 获取不限制的小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func (c *Client) GetUnlimitedQRCode(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetUnlimitedQRCodeResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetUnlimitedQRCodeResponse
	request, err := c.request(ctx, "wxa/getwxacodeunlimit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return newGetUnlimitedQRCodeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetUnlimitedQRCodeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45009:
		return "调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成"
	case 41030:
		return "page 不合法（页面不存在或者小程序没有发布、根路径前加 /或者携带参数）"
	case 40097:
		return "env_version 不合法"
	default:
		return resp.Result.Errmsg
	}
}
