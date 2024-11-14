package wechatqy

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinGetCallBackIpResponse struct {
	IpList  []string `json:"ip_list,omitempty"`
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
}

type CgiBinGetCallBackIpResult struct {
	Result CgiBinGetCallBackIpResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func NewCgiBinGetCallBackIpResult(result CgiBinGetCallBackIpResponse, body []byte, http gorequest.Response) *CgiBinGetCallBackIpResult {
	return &CgiBinGetCallBackIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetCallBackIp 获取企业微信回调IP段
// https://developer.work.weixin.qq.com/document/path/98988
func (c *Client) CgiBinGetCallBackIp(ctx context.Context, accessToken string, notMustParams ...*gorequest.Params) (*CgiBinGetCallBackIpResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinGetCallBackIpResponse
	request, err := c.request(ctx, apiUrl+"cgi-bin/getcallbackip?access_token="+accessToken, params, http.MethodGet, &response)
	return NewCgiBinGetCallBackIpResult(response, request.ResponseBody, request), err
}
