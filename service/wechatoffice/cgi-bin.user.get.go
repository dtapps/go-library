package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
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
}

func newCgiBinUserGetResult(result CgiBinUserGetResponse, body []byte, http gorequest.Response) *CgiBinUserGetResult {
	return &CgiBinUserGetResult{Result: result, Body: body, Http: http}
}

// CgiBinUserGet 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (c *Client) CgiBinUserGet(ctx context.Context, nextOpenid string, notMustParams ...*gorequest.Params) (*CgiBinUserGetResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/user/get?access_token=%s&next_openid=%s", c.getAccessToken(ctx), nextOpenid), params, http.MethodGet)
	if err != nil {
		return newCgiBinUserGetResult(CgiBinUserGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinUserGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinUserGetResult(response, request.ResponseBody, request), err
}
