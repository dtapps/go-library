package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gorequest"
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
// orderid 用户订单号	用户提交订单号
func (c *Client) CheckCost(ctx context.Context, orderid string, notMustParams ...*gorequest.Params) (*CheckCostResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	params.Set("orderid", orderid)      // 用户订单号	用户提交订单号

	// 响应
	var response CheckCostResponse

	// 请求
	request, err := c.requestXml(ctx, "checkCost.do", params, http.MethodGet, &response)
	return newCheckCostResult(response, request.ResponseBody, request), err
}
