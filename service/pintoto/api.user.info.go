package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiUserInfoResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Nickname     string  `json:"nickname"`      // 用户昵称
		Mobile       int64   `json:"mobile"`        // 注册号码
		Balance      float64 `json:"balance"`       // 账户余额
		FreezeAmount float64 `json:"freeze_amount"` // 冻结金额
	} `json:"data"`
	Code int `json:"code"`
}

type ApiUserInfoResult struct {
	Result ApiUserInfoResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newApiUserInfoResult(result ApiUserInfoResponse, body []byte, http gorequest.Response) *ApiUserInfoResult {
	return &ApiUserInfoResult{Result: result, Body: body, Http: http}
}

// ApiUserInfo 账号信息查询 https://www.showdoc.com.cn/1154868044931571/6269224958928211
func (c *Client) ApiUserInfo(ctx context.Context, notMustParams ...gorequest.Params) (*ApiUserInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/user/info", params)
	if err != nil {
		return newApiUserInfoResult(ApiUserInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiUserInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiUserInfoResult(response, request.ResponseBody, request), err
}
