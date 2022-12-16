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
	Error   int      `xml:"error"`   // 错误提示
}

type CheckCostResult struct {
	Result CheckCostResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newCheckCostResult(result CheckCostResponse, body []byte, http gorequest.Response, err error) *CheckCostResult {
	return &CheckCostResult{Result: result, Body: body, Http: http, Err: err}
}

// CheckCost 会员订单成本价查询接口
func (c *Client) CheckCost(ctx context.Context, orderId string) *CheckCostResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("orderid", orderId)
	params := gorequest.NewParamsWith(param)
	// 签名
	c.config.signStr = fmt.Sprintf("userid%vpwd%vorderid%v", c.GetUserId(), c.GetPwd(), orderId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/checkCost.do", params, http.MethodGet)
	// 定义
	var response CheckCostResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newCheckCostResult(response, request.ResponseBody, request, err)
}
