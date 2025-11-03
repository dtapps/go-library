package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferTransferBillsOutBillNoResponse struct {
	MchID          string `json:"mch_id"`                // 【商户号】 微信支付分配的商户号
	OutBillNo      string `json:"out_bill_no"`           // 【商户单号】 商户系统内部的商家单号，要求此参数只能由数字、大小写字母组成，在商户系统内部唯一
	TransferBillNo string `json:"transfer_bill_no"`      // 【商家转账订单号】 商家转账订单的主键，唯一定义此资源的标识
	AppID          string `json:"appid"`                 // 【商户AppID】 微信开放平台和微信公众平台为开发者的应用程序提供的唯一标识
	State          string `json:"state"`                 // 【单据状态】 ACCEPTED:转账已受理 PROCESSING:转账锁定资金中 WAIT_USER_CONFIRM:待收款用户确认 TRANSFERING:转账中 SUCCESS:转账成功 FAIL:转账失败 CANCELING:转账撤销中 CANCELLED:转账撤销完成
	TransferAmount int64  `json:"transfer_amount"`       // 【转账金额】 转账金额单位为"分"
	TransferRemark string `json:"transfer_remark"`       // 【转账备注】 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 【失败原因】 订单已失败或者已退资金时，会返回订单失败原因
	OpenID         string `json:"openid,omitempty"`      // 【收款用户OpenID】 用户在商户appid下的唯一标识
	UserName       string `json:"user_name,omitempty"`   // 【收款用户姓名】 收款方真实姓名，字段解密参考 如何使用API证书解密敏感字段
	CreateTime     string `json:"create_time"`           // 【单据创建时间】 单据受理成功时返回，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime     string `json:"update_time"`           // 【最后一次状态变更时间】 单据最后更新时间，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
}

// FundAppMchTransferTransferBillsOutBillNo 商户单号查询转账单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716437
func (c *Client) FundAppMchTransferTransferBillsOutBillNo(ctx context.Context, out_bill_no string, notMustParams ...*gorequest.Params) (response FundAppMchTransferTransferBillsOutBillNoResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/"+out_bill_no, params, http.MethodGet, &response, &apiError)
	return
}
