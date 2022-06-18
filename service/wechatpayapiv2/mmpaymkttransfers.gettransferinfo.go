package wechatpayapiv2

import (
	"encoding/xml"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gorequest"
)

type TransfersQueryResponse struct {
	ReturnCode     string `json:"return_code" xml:"return_code"`                         // 返回状态码
	ReturnMsg      string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`       // 返回信息
	ResultCode     string `json:"result_code" xml:"result_code"`                         // 业务结果
	ErrCode        string `json:"err_code,omitempty" xml:"err_code,omitempty"`           // 错误代码
	ErrCodeDes     string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"`   // 错误代码描述
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`               // 商户单号
	Appid          string `json:"appid" xml:"appid"`                                     // Appid
	MchId          string `json:"mch_id" xml:"mch_id"`                                   // 商户号
	DetailId       string `json:"detail_id" xml:"detail_id"`                             // 付款单号
	Status         string `json:"status" xml:"status"`                                   // 转账状态
	Reason         string `json:"reason,omitempty" xml:"reason,omitempty"`               // 失败原因
	Openid         string `json:"openid" xml:"openid"`                                   // 收款用户openid
	TransferName   string `json:"transfer_name,omitempty" xml:"transfer_name,omitempty"` // 收款用户姓名
	PaymentAmount  string `json:"payment_amount" xml:"payment_amount"`                   // 付款金额
	TransferTime   string `json:"transfer_time" xml:"transfer_time"`                     // 转账时间
	PaymentTime    string `json:"payment_time" xml:"payment_time"`                       // 付款成功时间
	Desc           string `json:"desc" xml:"desc"`                                       // 付款备注
}

type TransfersQueryResult struct {
	Result TransfersQueryResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func NewTransfersQueryResult(result TransfersQueryResponse, body []byte, http gorequest.Response, err error) *TransfersQueryResult {
	return &TransfersQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// TransfersQuery
// 付款到零钱 - 查询付款
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
func (app *App) TransfersQuery(partnerTradeNo string) *TransfersQueryResult {
	cert, err := app.P12ToPem()
	// 参数
	params := NewParams()
	params.Set("appid", app.appId)
	params.Set("mch_id", app.mchId)
	params.Set("nonce_str", gorandom.Alphanumeric(32))
	params.Set("partner_trade_no", partnerTradeNo)
	// 签名
	params.Set("sign", app.getMd5Sign(params))
	// 	请求
	request, err := app.request("https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo", params, cert)
	// 定义
	var response TransfersQueryResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return NewTransfersQueryResult(response, request.ResponseBody, request, err)
}
