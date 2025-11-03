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

	// 请求
	err = c.request(ctx, "/v3/fund-app/mch-transfer/transfer-bills", params, http.MethodPost, &response, &apiError)
	return
}
