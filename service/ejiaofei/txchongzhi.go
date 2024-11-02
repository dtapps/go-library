package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type TxChOngZhiParams struct {
	OrderID   string `json:"orderid"`   // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	Account   string `json:"account"`   // QQ号	需要充值的QQ号
	ProductID int64  `json:"productid"` // 产品id
	Amount    int64  `json:"amount"`    // 购买数量
	Ip        string `json:"ip"`        // 充值QQ号ip
	Times     string `json:"times"`     // 时间戳 格式：yyyyMMddhhmmss
}

type TxChOngZhiResponse struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 用户编号
	PorderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的QQ号
	ProductID int64    `xml:"productid"` // 充值产品id
	Amount    int64    `xml:"amount"`    // 购买数量
	State     int64    `xml:"state"`     // 订单状态
	StartTime string   `xml:"starttime"` // 开始时间
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

type TxChOngZhiResult struct {
	Result TxChOngZhiResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newTxChOngZhiResult(result TxChOngZhiResponse, body []byte, http gorequest.Response) *TxChOngZhiResult {
	return &TxChOngZhiResult{Result: result, Body: body, Http: http}
}

// TxChOngZhi 流量充值接口
// orderid = 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// account = QQ号 需要充值的QQ号
// productid = 产品id 可以通过2.5查询
// amount = 购买数量
// ip = 可以为空
func (c *Client) TxChOngZhi(ctx context.Context, orderid string, account string, productid int64, amount int64, notMustParams ...gorequest.Params) (*TxChOngZhiResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	params.Set("orderid", orderid)      // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	params.Set("account", account)      // QQ号 需要充值的QQ号
	params.Set("productid", productid)  // 产品id 可以通过2.5查询
	params.Set("amount", amount)        // 购买数量

	// 响应
	var response TxChOngZhiResponse

	// 请求
	request, err := c.requestXml(ctx, "txchongzhi.do", params, http.MethodGet, &response)
	return newTxChOngZhiResult(response, request.ResponseBody, request), err
}
