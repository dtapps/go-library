package ejiaofei

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type FindOrderResponse struct {
	Code      int64   `json:"code"`      // 返回状态编码
	State     int64   `json:"state"`     // 订单状态
	AppId     string  `json:"appId"`     // 用户编号
	OrderId   string  `json:"orderId"`   // 用户提交的订单号
	POrderId  string  `json:"pOrderId"`  // 平台订单号
	UserPrice float64 `json:"userPrice"` // 扣款金额
	EndTime   int64   `json:"endTime"`   // 结束时间
}

type FindOrderResult struct {
	Result FindOrderResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newFindOrderResult(result FindOrderResponse, body []byte, http gorequest.Response) *FindOrderResult {
	return &FindOrderResult{Result: result, Body: body, Http: http}
}

// FindOrder 订单查询接口
// orderId = 用户提交的订单号	是	用户提交的订单号，最长32位（用户保证其唯一性）
func (c *Client) FindOrder(ctx context.Context, orderID string, notMustParams ...gorequest.Params) (*FindOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appId", c.GetUserId())  // 用户编号 由鼎信商务提供
	params.Set("appSecret", c.GetPwd()) // 加密密码 由鼎信商务提供
	params.Set("orderId", orderID)      // 用户提交的订单号  用户提交的订单号，最长32位（用户保证其唯一性）
	// 请求
	request, err := c.requestJson(ctx, apiUrl+"/findOrder.do", params, http.MethodGet)
	if err != nil {
		return newFindOrderResult(FindOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response FindOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newFindOrderResult(response, request.ResponseBody, request), err
}
