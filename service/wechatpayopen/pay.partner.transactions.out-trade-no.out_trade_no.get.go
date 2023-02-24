package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type PayPartnerOutTradeNoOutTradeNoGetResponse struct {
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
}

type PayPartnerOutTradeNoOutTradeNoGetResult struct {
	Result   PayPartnerOutTradeNoOutTradeNoGetResponse // 结果
	Body     []byte                                    // 内容
	Http     gorequest.Response                        // 请求
	Err      error                                     // 错误
	ApiError ApiError                                  // 接口错误
}

func newPayPartnerOutTradeNoOutTradeNoGetResult(result PayPartnerOutTradeNoOutTradeNoGetResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerOutTradeNoOutTradeNoGetResult {
	return &PayPartnerOutTradeNoOutTradeNoGetResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerOutTradeNoOutTradeNoGet 商户订单号查询
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_2.shtml
func (c *Client) PayPartnerOutTradeNoOutTradeNoGet(ctx context.Context, out_trade_no string, notMustParams ...gorequest.Params) *PayPartnerOutTradeNoOutTradeNoGetResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/out-trade-no/"+out_trade_no, params, http.MethodGet)
	if err != nil {
		return newPayPartnerOutTradeNoOutTradeNoGetResult(PayPartnerOutTradeNoOutTradeNoGetResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response PayPartnerOutTradeNoOutTradeNoGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerOutTradeNoOutTradeNoGetResult(response, request.ResponseBody, request, err, apiError)
}
