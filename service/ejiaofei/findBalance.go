package ejiaofei

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
func (c *Client) FindBalance(ctx context.Context, notMustParams ...gorequest.Params) (*FindBalanceResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appId", c.GetUserId())  // 用户编号 由鼎信商务提供
	params.Set("appSecret", c.GetPwd()) // 加密密码 由鼎信商务提供
	// 请求
	request, err := c.requestJson(ctx, apiUrl+"/findBalance.do", params, http.MethodGet)
	if err != nil {
		return newFindBalanceResult(FindBalanceResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response FindBalanceResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newFindBalanceResult(response, request.ResponseBody, request), err
}
