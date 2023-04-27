package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiStartPushTicketResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type CgiBinComponentApiStartPushTicketResult struct {
	Result CgiBinComponentApiStartPushTicketResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newCgiBinComponentApiStartPushTicketResult(result CgiBinComponentApiStartPushTicketResponse, body []byte, http gorequest.Response) *CgiBinComponentApiStartPushTicketResult {
	return &CgiBinComponentApiStartPushTicketResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentApiStartPushTicket 启动ticket推送服务
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket_service.html
func (c *Client) CgiBinComponentApiStartPushTicket(ctx context.Context, notMustParams ...gorequest.Params) (*CgiBinComponentApiStartPushTicketResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId(ctx))      // 第三方平台appid
	params.Set("component_secret", c.GetComponentAppSecret(ctx)) // 第三方平台app_secret
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/api_start_push_ticket", params, http.MethodPost)
	if err != nil {
		return newCgiBinComponentApiStartPushTicketResult(CgiBinComponentApiStartPushTicketResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinComponentApiStartPushTicketResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentApiStartPushTicketResult(response, request.ResponseBody, request), err
}
