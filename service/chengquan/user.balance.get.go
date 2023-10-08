package chengquan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type UserBalanceGetResponse struct {
	Code    int    `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误信息
	Data    struct {
		AppID   string  `json:"app_id"`  // 商户账号
		Balance float64 `json:"balance"` // 商户余额(单位：元)
	} `json:"data"`
}

type UserBalanceGetResult struct {
	Result UserBalanceGetResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newUserBalanceGetResult(result UserBalanceGetResponse, body []byte, http gorequest.Response) *UserBalanceGetResult {
	return &UserBalanceGetResult{Result: result, Body: body, Http: http}
}

// UserBalanceGet 账号余额查询接口
// https://chengquan.cn/basicData/queryBalance.html
func (c *Client) UserBalanceGet(ctx context.Context, notMustParams ...gorequest.Params) (*UserBalanceGetResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/user/balance/get", params, http.MethodPost)
	if err != nil {
		return newUserBalanceGetResult(UserBalanceGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response UserBalanceGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newUserBalanceGetResult(response, request.ResponseBody, request), err
}
