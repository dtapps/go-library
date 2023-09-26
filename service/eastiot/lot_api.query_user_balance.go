package eastiot

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IotApiQueryUserBalanceResponse struct {
	Code int64 `json:"code"`
	Data struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryUserBalanceResult struct {
	Result IotApiQueryUserBalanceResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newIotApiQueryUserBalanceResult(result IotApiQueryUserBalanceResponse, body []byte, http gorequest.Response) *IotApiQueryUserBalanceResult {
	return &IotApiQueryUserBalanceResult{Result: result, Body: body, Http: http}
}

// IotApiQueryUserBalance 余额查询
// https://www.showdoc.com.cn/916774523755909/4857910459512420
func (c *Client) IotApiQueryUserBalance(ctx context.Context, notMustParams ...gorequest.Params) (*IotApiQueryUserBalanceResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/Api/IotApi/queryUserBalance", params, http.MethodPost)
	if err != nil {
		return newIotApiQueryUserBalanceResult(IotApiQueryUserBalanceResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response IotApiQueryUserBalanceResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIotApiQueryUserBalanceResult(response, request.ResponseBody, request), err
}
