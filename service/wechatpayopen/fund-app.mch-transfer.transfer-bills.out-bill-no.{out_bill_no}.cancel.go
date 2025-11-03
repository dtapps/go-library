package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferTransferBillsOutBillNoCancelResponse struct {
	OutBillNo      string `json:"out_bill_no"`      // 【商户单号】 商户系统内部的商家单号，要求此参数只能由数字、大小写字母组成，在商户系统内部唯一
	TransferBillNo string `json:"transfer_bill_no"` // 【微信转账单号】 微信转账单号，微信商家转账系统返回的唯一标识
	State          string `json:"state"`            // 【单据状态】 CANCELING: 撤销中；CANCELLED:已撤销
	UpdateTime     string `json:"update_time"`      // 【最后一次单据状态变更时间】 按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
}

// FundAppMchTransferTransferBillsOutBillNoCancel 撤销转账
// https://pay.weixin.qq.com/doc/v3/merchant/4012716458
func (c *Client) FundAppMchTransferTransferBillsOutBillNoCancel(ctx context.Context, out_bill_no string, notMustParams ...*gorequest.Params) (response FundAppMchTransferTransferBillsOutBillNoCancelResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/"+out_bill_no+"/cancel", params, http.MethodPost, &response, &apiError)
	return
}

// 【单据状态】
func (resp FundAppMchTransferTransferBillsOutBillNoCancelResponse) GetStateDesc() string {
	switch resp.State {
	case "CANCELING":
		return "转账撤销中，商户撤销请求受理成功，该笔转账正在撤销中，需查单确认撤销的转账单据状态（非终态）"
	case "CANCELLED":
		return "转账撤销完成，代表转账单据已撤销成功（终态）"
	default:
		return "未知状态"
	}
}
