package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinGetCallBackIpResponse struct {
	IpList []string `json:"ip_list,omitempty"`
}

type CgiBinGetCallBackIpResult struct {
	Result CgiBinGetCallBackIpResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func NewCgiBinGetCallBackIpResult(result CgiBinGetCallBackIpResponse, body []byte, http gorequest.Response) *CgiBinGetCallBackIpResult {
	return &CgiBinGetCallBackIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetCallBackIp 获取微信callback IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetCallBackIp(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*CgiBinGetCallBackIpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/get_api_domain_ip?access_token="+authorizerAccessToken, params, http.MethodGet)
	if err != nil {
		return NewCgiBinGetCallBackIpResult(CgiBinGetCallBackIpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinGetCallBackIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinGetCallBackIpResult(response, request.ResponseBody, request), err
}
