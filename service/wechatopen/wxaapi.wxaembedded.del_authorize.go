package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaApiWxaembeddedDelEmbeddedResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
}

type WxaApiWxaembeddedDelEmbeddedResult struct {
	Result WxaApiWxaembeddedDelEmbeddedResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newWxaApiWxaembeddedDelEmbeddedResult(result WxaApiWxaembeddedDelEmbeddedResponse, body []byte, http gorequest.Response) *WxaApiWxaembeddedDelEmbeddedResult {
	return &WxaApiWxaembeddedDelEmbeddedResult{Result: result, Body: body, Http: http}
}

// WxaApiWxaembeddedDelEmbedded 删除半屏小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/deleteEmbedded.html
func (c *Client) WxaApiWxaembeddedDelEmbedded(ctx context.Context, notMustParams ...gorequest.Params) (*WxaApiWxaembeddedDelEmbeddedResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxaapi/wxaembedded/del_embedded?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaApiWxaembeddedDelEmbeddedResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaApiWxaembeddedDelEmbeddedResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaApiWxaembeddedDelEmbeddedResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 89408:
		return "半屏小程序系统错误"
	case 89415:
		return "删除半屏小程序 appid 参数为空"
	case 89421:
		return "删除数据未找到"
	case 89422:
		return "删除状态异常"
	case 89431:
		return "不支持此类型小程序"
	case 89432:
		return "不是小程序"
	}
	return "系统繁忙"
}
