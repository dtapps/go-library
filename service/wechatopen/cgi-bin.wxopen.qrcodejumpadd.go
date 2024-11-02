package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpAddResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinWxOpenQrCodeJumpAddResult struct {
	Result CgiBinWxOpenQrCodeJumpAddResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newCgiBinWxOpenQrCodeJumpAddResult(result CgiBinWxOpenQrCodeJumpAddResponse, body []byte, http gorequest.Response) *CgiBinWxOpenQrCodeJumpAddResult {
	return &CgiBinWxOpenQrCodeJumpAddResult{Result: result, Body: body, Http: http}
}

// CgiBinWxOpenQrCodeJumpAdd 增加或修改二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpadd.html
func (c *Client) CgiBinWxOpenQrCodeJumpAdd(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*CgiBinWxOpenQrCodeJumpAddResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinWxOpenQrCodeJumpAddResponse
	request, err := c.request(ctx, "cgi-bin/wxopen/qrcodejumpadd?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newCgiBinWxOpenQrCodeJumpAddResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *CgiBinWxOpenQrCodeJumpAddResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 44990:
		return "接口请求太快（超过5次/秒）"
	case 85066:
		return "链接错误"
	case 85068:
		return "测试链接不是子链接"
	case 85069:
		return "校验文件失败"
	case 85070:
		return "URL命中黑名单，无法添加"
	case 85071:
		return "已添加该链接，请勿重复添加"
	case 85072:
		return "该链接已被占用"
	case 85073:
		return "二维码规则已满"
	case 85075:
		return "个人类型小程序无法设置二维码规则"
	default:
		return resp.Result.Errmsg
	}
}
