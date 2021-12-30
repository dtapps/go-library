package ejiaofei

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type GprsChOngZhiAdvanceParams struct {
	OrderID    string `json:"orderid"`    // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	Account    string `json:"account"`    // 充值手机号 需要充值的手机号
	Gprs       int    `json:"gprs"`       // 充值流量值 单位：MB（具体流量值请咨询商务）
	Area       int    `json:"area"`       // 充值流量范围 0 全国流量，1 省内流量
	EffectTime int    `json:"effecttime"` // 生效日期 0 即时生效，1次日生效，2 次月生效
	Validity   int    `json:"validity"`   // 流量有效期 传入月数，0为当月有效
	Times      string `json:"times"`      // 时间戳 格式：yyyyMMddhhmmss
}

type GprsChOngZhiAdvanceResponse struct {
	XMLName    xml.Name `xml:"response"`
	UserID     string   `xml:"userid"`     // 会员账号
	OrderID    string   `xml:"orderid"`    // 会员提交订单号
	PorderID   string   `xml:"Porderid"`   // 平台订单号
	Account    string   `xml:"account"`    // 充值手机号
	State      int      `xml:"state"`      // 订单状态
	StartTime  string   `xml:"starttime"`  // 开始时间
	EndTime    string   `xml:"endtime"`    // 结束时间
	Error      string   `xml:"error"`      // 错误提示
	UserPrice  float64  `xml:"userprice"`  // 会员购买价格
	Gprs       string   `xml:"gprs"`       // 充值流量值（单位MB）
	Area       string   `xml:"area"`       // 流量范围（0 全国流量，1省内流量）
	EffectTime string   `xml:"effecttime"` // 生效日期（0即时，1次日，2次月）
	Validity   string   `xml:"validity"`   // 流量有效期（显示月数，0为当月）
}

type GprsChOngZhiAdvanceResult struct {
	Result GprsChOngZhiAdvanceResponse // 结果
	Body   []byte                      // 内容
	Err    error                       // 错误
}

func NewGprsChOngZhiAdvanceResult(result GprsChOngZhiAdvanceResponse, body []byte, err error) *GprsChOngZhiAdvanceResult {
	return &GprsChOngZhiAdvanceResult{Result: result, Body: body, Err: err}
}

// GprsChOngZhiAdvance 流量充值接口
func (app *App) GprsChOngZhiAdvance(param GprsChOngZhiAdvanceParams) *GprsChOngZhiAdvanceResult {
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vgprs%varea%veffecttime%vvalidity%vtimes%v", app.UserID, app.Pwd, param.OrderID, param.Account, param.Gprs, param.Area, param.EffectTime, param.Validity, param.Times)
	// 请求
	b, _ := json.Marshal(&param)
	var params map[string]interface{}
	_ = json.Unmarshal(b, &params)
	body, err := app.request("http://api.ejiaofei.net:11140/gprsChongzhiAdvance.do", params, http.MethodGet)
	// 定义
	var response GprsChOngZhiAdvanceResponse
	err = xml.Unmarshal(body, &response)
	return NewGprsChOngZhiAdvanceResult(response, body, err)
}
