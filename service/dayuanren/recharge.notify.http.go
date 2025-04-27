package dayuanren

import (
	"context"
	"encoding/json"
	"net/http"
)

type ResponseRechargeNotifyHttp struct {
	Userid       string `json:"userid,omitempty"`        // 商户ID
	OrderNumber  string `json:"order_number,omitempty"`  // true
	OutTradeNum  string `json:"out_trade_num,omitempty"` // 商户订单号
	Otime        string `json:"otime,omitempty"`         // 成功/失败时间，10位时间戳
	State        string `json:"state,omitempty"`         // 充值状态；-1取消， 0充值中， 1充值成功 ，2充值失败，3部分成功（-1,2做失败处理；1做成功处理；3做部分成功处理）
	Mobile       string `json:"mobile,omitempty"`        // 充值手机号
	Remark       string `json:"remark,omitempty"`        // 备注信息
	ChargeAmount string `json:"charge_amount,omitempty"` // 充值成功面额
	Voucher      string `json:"voucher,omitempty"`       // 凭证
	ChargeKami   string `json:"charge_kami,omitempty"`   // 卡密/流水号
	Sign         string `json:"sign,omitempty"`          // 签名字符串，用于验签,以保证回调可靠性。
}

// RechargeNotifyHttp 充值结果通知-异步通知
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097255
func (c *Client) RechargeNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateJson ResponseRechargeNotifyHttp, err error) {
	err = json.NewDecoder(r.Body).Decode(&validateJson)
	return validateJson, err
}
