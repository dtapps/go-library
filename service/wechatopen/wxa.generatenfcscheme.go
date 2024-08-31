package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GenerateNFCSchemeResponse struct {
	Errcode  int    `json:"errcode"`  // 错误码
	Errmsg   string `json:"errmsg"`   // 错误信息
	Openlink string `json:"openlink"` // 生成的小程序 scheme 码
}

type GenerateNFCSchemeResult struct {
	Result GenerateNFCSchemeResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newGenerateNFCSchemeResult(result GenerateNFCSchemeResponse, body []byte, http gorequest.Response) *GenerateNFCSchemeResult {
	return &GenerateNFCSchemeResult{Result: result, Body: body, Http: http}
}

// GenerateNFCScheme 获取 NFC 的小程序 scheme
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateNFCScheme.html
func (c *Client) GenerateNFCScheme(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GenerateNFCSchemeResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/generatenfcscheme")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GenerateNFCSchemeResponse
	request, err := c.request(ctx, span, "wxa/generatenfcscheme?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	return newGenerateNFCSchemeResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GenerateNFCSchemeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40002:
		return "暂无生成权限（个人主体小程序无权限，或者NFC 能力的小程序未申请权限）"
	case 40013:
		return "生成权限被封禁"
	case 85079:
		return "小程序没有线上版本，即小程序尚未发布，不可进行该操作"
	case 40165:
		return "参数path填写错误，更正后重试"
	case 40212:
		return "参数query填写错误 ，query格式遵循URL标准，即k1=v1&k2=v2"
	case 85402:
		return "参数env_version填写错误，更正后重试"
	case 44990:
		return "频率过快，超过100次/秒；降低调用频率"
	case 44993:
		return "单天生成加密 URL Scheme+URL Link 数量超过上限50万"
	case 9800003:
		return "model_id检查不通过"
	case 9800007:
		return "此model_id尚未获得该能力，请能力申请通过后再试"
	case 9800008:
		return "能力类型为一机一码，sn不能为空"
	case 9800009:
		return "能力类型为一型一码，sn需为空"
	default:
		return resp.Result.Errmsg
	}
}
