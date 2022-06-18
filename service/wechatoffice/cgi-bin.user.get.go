package wechatoffice

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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

func NewCgiBinUserGetResult(result CgiBinUserGetResponse, body []byte, http gorequest.Response, err error) *CgiBinUserGetResult {
	return &CgiBinUserGetResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinUserGet 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (c *Client) CgiBinUserGet(nextOpenid string) *CgiBinUserGetResult {
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", c.getAccessToken(), nextOpenid), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinUserGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinUserGetResult(response, request.ResponseBody, request, err)
}
