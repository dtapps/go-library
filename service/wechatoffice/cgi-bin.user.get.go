package wechatoffice

import (
	"encoding/json"
	"fmt"
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
	Err    error                 // 错误
}

func NewCgiBinUserGetResult(result CgiBinUserGetResponse, body []byte, err error) *CgiBinUserGetResult {
	return &CgiBinUserGetResult{Result: result, Body: body, Err: err}
}

// CgiBinUserGet 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (app *App) CgiBinUserGet(nextOpenid string) *CgiBinUserGetResult {
	app.AccessToken = app.GetAccessToken()
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", app.AccessToken, nextOpenid), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinUserGetResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinUserGetResult(response, body, err)
}
