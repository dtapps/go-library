package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gorequest"
)

type MmpaymkttransfersGethbinfoResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code" xml:"result_code"`                       // 业务结果
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	MchBillno    string `json:"mch_billno" xml:"mch_billno"`       // 商户订单号
	MchId        string `json:"mch_id" xml:"mch_id"`               // 商户号
	DetailId     string `json:"detail_id" xml:"detail_id"`         // 红包单号
	Status       string `json:"status" xml:"status"`               // 红包状态
	SendType     string `json:"send_type" xml:"send_type"`         // 发放类型
	HbType       string `json:"hb_type" xml:"hb_type"`             // 红包类型
	TotalNum     int64  `json:"total_num" xml:"total_num"`         // 红包个数
	TotalAmount  int64  `json:"total_amount" xml:"total_amount"`   // 红包金额
	Reason       string `json:"reason" xml:"reason"`               // 失败原因
	SendTime     string `json:"send_time" xml:"send_time"`         // 红包发送时间
	RefundTime   string `json:"refund_time" xml:"refund_time"`     // 红包退款时间
	RefundAmount int64  `json:"refund_amount" xml:"refund_amount"` // 红包退款金额
	Wishing      string `json:"wishing" xml:"wishing"`             // 祝福语
	Remark       string `json:"remark" xml:"remark"`               // 活动描述
	ActName      string `json:"act_name" xml:"act_name"`           // 活动名称
	Hblist       []struct {
		Openid   string `json:"openid" xml:"openid"`     // 领取红包的Openid
		Amount   int64  `json:"amount" xml:"amount"`     // 金额
		Rcv_time string `json:"rcv_time" xml:"rcv_time"` // 接收时间
	} `json:"hblist" xml:"hblist"` // 裂变红包领取列表
}

type MmpaymkttransfersGethbinfoResult struct {
	Result MmpaymkttransfersGethbinfoResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
}

func newMmpaymkttransfersGethbinfoResult(result MmpaymkttransfersGethbinfoResponse, body []byte, http gorequest.Response) *MmpaymkttransfersGethbinfoResult {
	return &MmpaymkttransfersGethbinfoResult{Result: result, Body: body, Http: http}
}

// MmpaymkttransfersGethbinfo
// 付款到零钱 - 查询红包记录
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon_sl.php?chapter=13_6&index=5
func (c *Client) MmpaymkttransfersGethbinfo(ctx context.Context, notMustParams ...gorequest.Params) (*MmpaymkttransfersGethbinfoResult, error) {
	cert, err := c.P12ToPem()
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串
	// 签名
	params.Set("sign", c.getMd5Sign(params))
	// 	请求
	request, err := c.request(ctx, apiUrl+"/mmpaymkttransfers/sendgroupredpack", params, true, cert)
	if err != nil {
		return newMmpaymkttransfersGethbinfoResult(MmpaymkttransfersGethbinfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MmpaymkttransfersGethbinfoResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newMmpaymkttransfersGethbinfoResult(response, request.ResponseBody, request), err
}
