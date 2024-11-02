package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CreateQRCodeResponse struct {
	Errcode     int    `json:"errcode"`               // 错误码
	Errmsg      string `json:"errmsg"`                // 错误信息
	ContentType string `json:"contentType,omitempty"` // 内容类型
	Buffer      any    `json:"buffer"`                // 图片 Buffer
}

type CreateQRCodeResult struct {
	Result CreateQRCodeResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newCreateQRCodeResult(result CreateQRCodeResponse, body []byte, http gorequest.Response) *CreateQRCodeResult {
	return &CreateQRCodeResult{Result: result, Body: body, Http: http}
}

// CreateQRCode 获取小程序二维码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
func (c *Client) CreateQRCode(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*CreateQRCodeResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CreateQRCodeResponse
	request, err := c.request(ctx, "cgi-bin/wxaapp/createwxaqrcode?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return newCreateQRCodeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *CreateQRCodeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45029:
		return "生成码个数总和到达最大个数限制"
	case 40159:
		return "path 不能为空，且长度不能大于 128 字节"
	case 85096:
		return "scancode_time为系统保留参数，不允许配置"
	default:
		return resp.Result.Errmsg
	}
}
