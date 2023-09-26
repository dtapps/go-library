package ejiaofei

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CheckCostResponse struct {
	XMLName xml.Name `xml:"response"`
	UserID  string   `xml:"userid"`  // 用户账号
	OrderID string   `xml:"orderid"` // 用户提交订单号
	Face    float64  `xml:"face"`    // 官方价格
	Price   float64  `xml:"price"`   // 用户成本价
	Error   int64    `xml:"error"`   // 错误提示
}

type CheckCostResult struct {
	Result CheckCostResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newCheckCostResult(result CheckCostResponse, body []byte, http gorequest.Response) *CheckCostResult {
	return &CheckCostResult{Result: result, Body: body, Http: http}
}

// CheckCost 会员订单成本价查询接口
// orderID 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
func (c *Client) CheckCost(ctx context.Context, notMustParams ...gorequest.Params) (*CheckCostResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 签名
	c.config.signStr = fmt.Sprintf("userid%vpwd%vorderid%v", c.GetUserId(), c.GetPwd(), params.Get("orderid"))
	// 请求
	request, err := c.request(ctx, apiUrl+"/checkCost.do", params, http.MethodGet)
	if err != nil {
		return newCheckCostResult(CheckCostResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CheckCostResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newCheckCostResult(response, request.ResponseBody, request), err
}
