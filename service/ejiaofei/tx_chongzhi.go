package ejiaofei

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type TxChOngZhiParams struct {
	OrderID   string `json:"orderid"`   // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	Account   string `json:"account"`   // QQ号	需要充值的QQ号
	ProductID int    `json:"productid"` // 产品id
	Amount    int    `json:"amount"`    // 购买数量
	Ip        string `json:"ip"`        // 充值QQ号ip
	Times     string `json:"times"`     // 时间戳 格式：yyyyMMddhhmmss
}

type TxChOngZhiResult struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 用户编号
	PorderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的QQ号
	ProductID int      `xml:"productid"` // 充值产品id
	Amount    int      `xml:"amount"`    // 购买数量
	State     int      `xml:"state"`     // 订单状态
	StartTime string   `xml:"starttime"` // 开始时间
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

// TxChOngZhi 流量充值接口
func (app *App) TxChOngZhi(param TxChOngZhiParams) (body []byte, err error) {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vproductid%vamount%vip%vtimes%v", app.UserID, app.Pwd, param.OrderID, param.Account, param.ProductID, param.Amount, param.Ip, param.Times)
	// 请求
	b, _ := json.Marshal(&param)
	var params map[string]interface{}
	_ = json.Unmarshal(b, &params)
	body, err = app.request("http://api.ejiaofei.net:11140/txchongzhi.do", params, http.MethodGet)
	return body, err
}
