package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferTransferBillsResponse struct {
	OutBillNo      string `json:"out_bill_no"`      // 【商户单号】 商户系统内部的商家单号，要求此参数只能由数字、大小写字母组成，在商户系统内部唯一
	TransferBillNo string `json:"transfer_bill_no"` // 【微信转账单号】 微信转账单号，微信商家转账系统返回的唯一标识
	CreateTime     string `json:"create_time"`      // 【单据创建时间】 单据受理成功时返回，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
	State          string `json:"state"`            // 【单据状态】 商家转账订单状态
	PackageInfo    string `json:"package_info"`     // 【跳转领取页面的package信息】 跳转微信支付收款页的package信息，APP调起用户确认收款或者JSAPI调起用户确认收款 时需要使用的参数。仅当转账单据状态为WAIT_USER_CONFIRM时返回。
}

// FundAppMchTransferTransferBills 发起转账
// https://pay.weixin.qq.com/doc/v3/merchant/4012716434
func (c *Client) FundAppMchTransferTransferBills(ctx context.Context, notMustParams ...*gorequest.Params) (response FundAppMchTransferTransferBillsResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetSpAppid()) // 商户AppID

	// 请求
	err = c.NewRequest(ctx, "/v3/fund-app/mch-transfer/transfer-bills", params, http.MethodPost, &response, &apiError)
	return
}

// 【单据状态】 商家转账订单状态
func (resp FundAppMchTransferTransferBillsResponse) GetStateDesc() string {
	switch resp.State {
	case "ACCEPTED":
		return "转账已受理，可原单重试（非终态）"
	case "PROCESSING":
		return "转账锁定资金中。如果一直停留在该状态，建议检查账户余额是否足够，如余额不足，可充值后再原单重试（非终态）"
	case "WAIT_USER_CONFIRM":
		return "待收款用户确认，当前转账单据资金已锁定，可拉起微信收款确认页面进行收款确认（非终态）"
	case "TRANSFERING":
		return "转账中，可拉起微信收款确认页面再次重试确认收款（非终态）"
	case "SUCCESS":
		return "转账成功，表示转账单据已成功（终态）"
	case "FAIL":
		return "转账失败，表示该笔转账单据已失败。若需重新向用户转账，请重新生成单据并再次发起（终态）"
	case "CANCELING":
		return "转账撤销中，商户撤销请求受理成功，该笔转账正在撤销中，需查单确认撤销的转账单据状态（非终态）"
	case "CANCELLED":
		return "转账撤销完成，代表转账单据已撤销成功（终态）"
	default:
		return "未知状态"
	}
}
