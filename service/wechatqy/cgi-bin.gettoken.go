package wechatqy

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinGetTokenResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type CgiBinGetTokenResult struct {
	Result CgiBinGetTokenResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newCgiBinGetTokenResult(result CgiBinGetTokenResponse, body []byte, http gorequest.Response) *CgiBinGetTokenResult {
	return &CgiBinGetTokenResult{Result: result, Body: body, Http: http}
}

// CgiBinGetToken 获取access_token
// https://open.work.weixin.qq.com/api/doc/90000/90135/91039
func (c *Client) CgiBinGetToken(ctx context.Context, notMustParams ...gorequest.Params) (*CgiBinGetTokenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/cgi-bin/gettoken?corpid=%s&corpsecret=%s", c.GetAppId(), c.GetSecret()), params, http.MethodGet)
	if err != nil {
		return newCgiBinGetTokenResult(CgiBinGetTokenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinGetTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinGetTokenResult(response, request.ResponseBody, request), err
}
