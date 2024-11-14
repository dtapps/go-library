package ejiaofei

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type FindBalanceResponse struct {
	Code    int64   `json:"code"`    // 返回状态编码
	Balance float64 `json:"balance"` // 用户余额
}

type FindBalanceResult struct {
	Result FindBalanceResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newFindBalanceResult(result FindBalanceResponse, body []byte, http gorequest.Response) *FindBalanceResult {
	return &FindBalanceResult{Result: result, Body: body, Http: http}
}

// FindBalance 余额查询接口
func (c *Client) FindBalance(ctx context.Context, notMustParams ...*gorequest.Params) (*FindBalanceResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appId", c.GetUserId())  // 用户编号 由鼎信商务提供
	params.Set("appSecret", c.GetPwd()) // 加密密码 由鼎信商务提供

	// 响应
	var response FindBalanceResponse

	// 请求
	request, err := c.requestJson(ctx, "findBalance.do", params, http.MethodGet, &response)
	return newFindBalanceResult(response, request.ResponseBody, request), err
}
