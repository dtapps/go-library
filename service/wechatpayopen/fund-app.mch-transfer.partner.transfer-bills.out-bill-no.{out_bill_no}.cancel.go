package wechatpayopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferPartnerTransferBillsOutBillNoCancelResponse struct {
	SubMchid       string `json:"sub_mchid"`             // 子商户号，微信支付分配的商户号，出资商户号
	Mchid          string `json:"mchid"`                 // 商户号，微信支付分配的商户号
	OutBillNo      string `json:"out_bill_no"`           // 商户单号，商户系统内部的商家单号，要求此参数只能由数字、大小写字母组成，在商户系统内部唯一
	TransferBillNo string `json:"transfer_bill_no"`      // 商家转账订单号，商家转账订单的主键，唯一定义此资源的标识
	Appid          string `json:"appid"`                 // 商户AppID，申请商户号的AppID或商户号绑定的AppID（企业号corpid即为此AppID）
	State          string `json:"state"`                 // 单据状态
	TransferAmount int64  `json:"transfer_amount"`       // 转账金额，转账金额单位为"分"
	TransferRemark string `json:"transfer_remark"`       // 转账备注，单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 失败原因，订单已失败或者已退资金时，返回失败原因
	Openid         string `json:"openid,omitempty"`      // 收款用户OpenID，商户AppID下，某用户的OpenID
	UserName       string `json:"user_name,omitempty"`   // 收款用户姓名，收款方真实姓名。支持标准RSA算法和国密算法，公钥由微信侧提供
	CreateTime     string `json:"create_time"`           // 单据创建时间，单据受理成功时返回，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime     string `json:"update_time"`           // 最后一次状态变更时间，单据最后更新时间，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
}

// FundAppMchTransferPartnerTransferBillsOutBillNoCancel 撤销转账
// https://pay.weixin.qq.com/doc/v3/partner/4015469118
func (c *Client) FundAppMchTransferPartnerTransferBillsOutBillNoCancel(ctx context.Context, out_bill_no string, notMustParams ...*gorequest.Params) (response FundAppMchTransferPartnerTransferBillsOutBillNoCancelResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, fmt.Sprintf("/v3/fund-app/mch-transfer/partner/transfer-bills/out-bill-no/%s/cancel", out_bill_no), params, http.MethodPost, &response, &apiError)
	return
}

// 【单据状态】
func (resp FundAppMchTransferPartnerTransferBillsOutBillNoCancelResponse) GetStateDesc() string {
	switch resp.State {
	case "ACCEPTED":
		return "转账已受理"
	case "PROCESSING":
		return "转账处理中，转账结果尚未明确，如一直处于此状态，建议检查账户余额是否足够"
	case "WAIT_USER_CONFIRM":
		return "待收款用户确认，可拉起微信收款确认页面进行收款确认"
	case "TRANSFERING":
		return "转账结果尚未明确，可拉起微信收款确认页面再次重试确认收款"
	case "SUCCESS":
		return "转账成功"
	case "FAIL":
		return "转账失败"
	case "CANCELING":
		return "商户撤销请求受理成功，该笔转账正在撤销中"
	case "CANCELLED":
		return "转账撤销完成"
	default:
		return "未知状态"
	}
}
