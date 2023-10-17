package ejiaofei

import (
	"context"
	"encoding/xml"
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
// orderid 用户订单号	用户提交订单号
func (c *Client) CheckCost(ctx context.Context, orderid string, notMustParams ...gorequest.Params) (*CheckCostResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	params.Set("orderid", orderid)      // 用户订单号	用户提交订单号
	// 请求
	request, err := c.requestXml(ctx, apiUrl+"/checkCost.do", params, http.MethodGet)
	if err != nil {
		return newCheckCostResult(CheckCostResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CheckCostResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newCheckCostResult(response, request.ResponseBody, request), err
}
