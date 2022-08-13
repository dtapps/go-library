package wechatoffice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinUserGetResponse struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type CgiBinUserGetResult struct {
	Result CgiBinUserGetResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newCgiBinUserGetResult(result CgiBinUserGetResponse, body []byte, http gorequest.Response, err error) *CgiBinUserGetResult {
	return &CgiBinUserGetResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinUserGet 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (c *Client) CgiBinUserGet(ctx context.Context, nextOpenid string) *CgiBinUserGetResult {
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/user/get?access_token=%s&next_openid=%s", c.getAccessToken(ctx), nextOpenid), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinUserGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinUserGetResult(response, request.ResponseBody, request, err)
}
