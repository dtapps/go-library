package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type FundAppMchTransferPartnerTransferBillsResponse struct {
	SubMchid       string `json:"sub_mchid"`        // 子商户号，微信支付分配的商户号，出资商户
	OutBillNo      string `json:"out_bill_no"`      // 商户单号，商户系统内部的商家单号，只能由数字、大小写字母组成，在商户系统内部唯一
	TransferBillNo string `json:"transfer_bill_no"` // 微信支付转账单号，转账单的唯一标识
	CreateTime     string `json:"create_time"`      // 受理时间，受理成功时会返回，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
	PackageInfo    string `json:"package_info"`     // 跳转领取页面的package信息，跳转微信支付收款页的package信息，APP调起用户确认收款或者JSAPI调起用户确认收款时需要使用的参数
}

// FundAppMchTransferPartnerTransferBills 发起转账
// https://pay.weixin.qq.com/doc/v3/partner/4015469096
func (c *Client) FundAppMchTransferPartnerTransferBills(ctx context.Context, notMustParams ...*gorequest.Params) (response FundAppMchTransferPartnerTransferBillsResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSubAppid())     // 商户AppID

	// 使用安全方法获取user_name参数
	if userNameStr, ok := params.GetStringOK("user_name"); ok && userNameStr != "" {
		userName, err := EncryptOAEPWithPublicKey(userNameStr, c.config.publicKey)
		if err != nil {
			return response, apiError, err
		}
		params.Set("user_name", userName)
	}

	// 请求
	err = c.NewRequest(ctx, "/v3/fund-app/mch-transfer/partner/transfer-bills", params, http.MethodPost, &response, &apiError)
	return
}
