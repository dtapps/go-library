package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type PayPartnerTransactionsIdTransactionIdGetResponse struct {
	SpAppid        string    `json:"sp_appid"`         // 服务商应用ID
	SpMchid        string    `json:"sp_mchid"`         // 服务商户号
	SubAppid       string    `json:"sub_appid"`        // 子商户应用ID
	SubMchid       string    `json:"sub_mchid"`        // 子商户号
	OutTradeNo     string    `json:"out_trade_no"`     // 商户订单号
	TransactionId  string    `json:"transaction_id"`   // 微信支付订单号
	TradeType      string    `json:"trade_type"`       // 交易类型
	TradeState     string    `json:"trade_state"`      // 交易状态
	TradeStateDesc string    `json:"trade_state_desc"` // 交易状态描述
	BankType       string    `json:"bank_type"`        // 付款银行
	Attach         string    `json:"attach"`           // 附加数据
	SuccessTime    time.Time `json:"success_time"`     // 支付完成时间
	Payer          struct {
		SpOpenid  string `json:"sp_openid"`  // 用户服务标识
		SubOpenid string `json:"sub_openid"` // 用户子标识
	} `json:"payer"` // 支付者
	Amount struct {
		Total         int    `json:"total"`          // 总金额
		PayerTotal    int    `json:"payer_total"`    // 用户支付金额
		Currency      string `json:"currency"`       // 货币类型
		PayerCurrency string `json:"payer_currency"` // 用户支付币种
	} `json:"amount"` // 订单金额
	SceneInfo struct {
		DeviceId string `json:"device_id"` // 商户端设备号
	} `json:"scene_info"` // 场景信息
	PromotionDetail []interface{} `json:"promotion_detail"` // 优惠功能
}

type PayPartnerTransactionsIdTransactionIdGetResult struct {
	Result   PayPartnerTransactionsIdTransactionIdGetResponse // 结果
	Body     []byte                                           // 内容
	Http     gorequest.Response                               // 请求
	Err      error                                            // 错误
	ApiError ApiError                                         // 接口错误
}

func newPayPartnerTransactionsIdTransactionIdGetResult(result PayPartnerTransactionsIdTransactionIdGetResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsIdTransactionIdGetResult {
	return &PayPartnerTransactionsIdTransactionIdGetResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsIdTransactionIdGet 微信支付订单号查询
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_2.shtml
func (c *Client) PayPartnerTransactionsIdTransactionIdGet(ctx context.Context, transactionId string, notMustParams ...*gorequest.Params) *PayPartnerTransactionsIdTransactionIdGetResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/id/"+transactionId, params, http.MethodGet)
	if err != nil {
		return newPayPartnerTransactionsIdTransactionIdGetResult(PayPartnerTransactionsIdTransactionIdGetResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response PayPartnerTransactionsIdTransactionIdGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsIdTransactionIdGetResult(response, request.ResponseBody, request, err, apiError)
}
