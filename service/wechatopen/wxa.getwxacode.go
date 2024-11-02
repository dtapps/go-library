package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetQRCodeResponse struct {
	Errcode     int    `json:"errcode"`               // 错误码
	Errmsg      string `json:"errmsg"`                // 错误信息
	ContentType string `json:"contentType,omitempty"` // 内容类型
	Buffer      any    `json:"buffer"`                // 图片 Buffer
}

type GetQRCodeResult struct {
	Result GetQRCodeResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGetQRCodeResult(result GetQRCodeResponse, body []byte, http gorequest.Response) *GetQRCodeResult {
	return &GetQRCodeResult{Result: result, Body: body, Http: http}
}

// GetQRCode 获取小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
func (c *Client) GetQRCode(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetQRCodeResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetQRCodeResponse
	request, err := c.request(ctx, "wxa/getwxacode?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return newGetQRCodeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetQRCodeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 40159:
		return "path 不能为空，且长度不能大于1024"
	case 45029:
		return "生成码个数总和到达最大个数限制"
	case 85096:
		return "scancode_time为系统保留参数，不允许配置"
	case 40097:
		return "参数错误"
	default:
		return resp.Result.Errmsg
	}
}
