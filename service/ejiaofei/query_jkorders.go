package ejiaofei

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type QueryJkOrdersResponse struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 会员账号
	POrderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的手机号码
	Face      string   `xml:"face"`      // 充值面值
	Amount    string   `xml:"amount"`    // 购买数量
	StartTime string   `xml:"starttime"` // 开始时间
	State     string   `xml:"state"`     // 订单状态
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

type QueryJkOrdersResult struct {
	Result QueryJkOrdersResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newQueryJkOrdersResult(result QueryJkOrdersResponse, body []byte, http gorequest.Response) *QueryJkOrdersResult {
	return &QueryJkOrdersResult{Result: result, Body: body, Http: http}
}

// QueryJkOrders 通用查询接口
// orderId 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
func (c *Client) QueryJkOrders(ctx context.Context, notMustParams ...gorequest.Params) (*QueryJkOrdersResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 签名
	c.config.signStr = fmt.Sprintf("userid%vpwd%vorderid%v", c.GetUserId(), c.GetPwd(), params.Get("orderid"))
	// 请求
	request, err := c.request(ctx, apiUrl+"/query_jkorders.do", params, http.MethodGet)
	if err != nil {
		return newQueryJkOrdersResult(QueryJkOrdersResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response QueryJkOrdersResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newQueryJkOrdersResult(response, request.ResponseBody, request), err
}
