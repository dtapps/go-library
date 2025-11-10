package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CreateQRCodeResponse struct {
	APIResponse        // 错误
	ContentType string `json:"contentType,omitempty"` // 内容类型
	Buffer      any    `json:"buffer"`                // 图片 Buffer
}

// CreateQRCode 获取小程序二维码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
func (c *Client) CreateQRCode(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response CreateQRCodeResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/wxaapp/createwxaqrcode?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return
}

// ErrcodeInfo 错误描述
func GetCreateQRCodeErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 45029:
		return "生成码个数总和到达最大个数限制"
	case 40159:
		return "path 不能为空，且长度不能大于 128 字节"
	case 85096:
		return "scancode_time为系统保留参数，不允许配置"
	default:
		return errmsg
	}
}
