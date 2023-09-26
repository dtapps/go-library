package ejiaofei

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
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
func (c *Client) TxChOngZhi(ctx context.Context, notMustParams ...gorequest.Params) (*TxChOngZhiResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 签名
	c.config.signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vproductid%vamount%vip%vtimes%v", c.GetUserId(), c.GetPwd(), params.Get("orderid"), params.Get("account"), params.Get("productid"), params.Get("amount"), params.Get("ip"), params.Get("times"))
	// 请求
	request, err := c.request(ctx, apiUrl+"/txchongzhi.do", params, http.MethodGet)
	if err != nil {
		return newTxChOngZhiResult(TxChOngZhiResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TxChOngZhiResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newTxChOngZhiResult(response, request.ResponseBody, request), err
}
