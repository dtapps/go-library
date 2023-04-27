package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func NewGetCallBackIpResult(result GetCallBackIpResponse, body []byte, http gorequest.Response) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Body: body, Http: http}
}

// CgiBinGetApiDomainIp 获取微信服务器IP地址
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (c *Client) CgiBinGetApiDomainIp(ctx context.Context, componentAccessToken string, notMustParams ...gorequest.Params) (*GetCallBackIpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/get_api_domain_ip?access_token="+componentAccessToken, params, http.MethodGet)
	if err != nil {
		return NewGetCallBackIpResult(GetCallBackIpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetCallBackIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return NewGetCallBackIpResult(response, request.ResponseBody, request), err
}
