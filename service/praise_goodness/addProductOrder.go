package praise_goodness

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type AddProductOrderResponse struct {
	Code int      `json:"code"` // 1：请求成功 -1：请求失败
	Msg  string   `json:"msg"`  // 返回说明
	Time string   `json:"time"` // 时间戳
	Data struct{} `json:"data"`
}

type AddProductOrderResult struct {
	Result AddProductOrderResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newAddProductOrderResult(result AddProductOrderResponse, body []byte, http gorequest.Response) *AddProductOrderResult {
	return &AddProductOrderResult{Result: result, Body: body, Http: http}
}

// AddProductOrder 下单接口
func (c *Client) AddProductOrder(ctx context.Context, notMustParams ...gorequest.Params) (*AddProductOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "api/order/addProductOrder", params, http.MethodPost)
	if err != nil {
		return newAddProductOrderResult(AddProductOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response AddProductOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newAddProductOrderResult(response, request.ResponseBody, request), err
}
