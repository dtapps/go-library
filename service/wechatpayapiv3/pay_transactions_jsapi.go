package wechatpayapiv3

import (
	"encoding/json"
	"time"
)

// PayTransactionsJsapi 入参
type PayTransactionsJsapi struct {
	Description string                          `json:"description"`           //【是】商品描述
	OutTradeNo  string                          `json:"out_trade_no"`          //【是】商户订单号
	TimeExpire  time.Time                       `json:"time_expire,omitempty"` //【否】交易结束时间
	Attach      string                          `json:"attach,omitempty"`      //【否】附加数据
	NotifyUrl   string                          `json:"notify_url"`            //【是】通知地址
	GoodsTag    string                          `json:"goods_tag,omitempty"`   //【否】订单优惠标记
	Amount      *PayTransactionsJsapiAmount     `json:"amount"`                //【是】订单金额
	Payer       *PayTransactionsJsapiPayer      `json:"payer"`                 //【是】支付者
	Detail      *PayTransactionsJsapiDetail     `json:"detail,omitempty"`      //【否】优惠功能
	SceneInfo   *PayTransactionsJsapiSceneInfo  `json:"scene_info,omitempty"`  //【否】场景信息
	SettleInfo  *PayTransactionsJsapiSettleInfo `json:"settle_info,omitempty"` //【否】结算信息
}

// PayTransactionsJsapiAmount 订单金额
type PayTransactionsJsapiAmount struct {
	Total    int    `json:"total"`              //【是】总金额
	Currency string `json:"currency,omitempty"` //【否】货币类型
}

// PayTransactionsJsapiPayer 支付者
type PayTransactionsJsapiPayer struct {
	Openid string `json:"openid"` //【是】用户标识
}

// PayTransactionsJsapiDetail 优惠功能
type PayTransactionsJsapiDetail struct {
	CostPrice   int                                     `json:"cost_price,omitempty"`   //【否】订单原价
	InvoiceId   string                                  `json:"invoice_id,omitempty"`   //【否】商品小票ID
	GoodsDetail []PayTransactionsJsapiDetailGoodsDetail `json:"goods_detail,omitempty"` //【否】单品列表
}

// PayTransactionsJsapiDetailGoodsDetail 单品列表
type PayTransactionsJsapiDetailGoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`            //【是】商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` //【否】微信侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         //【否】商品名称
	Quantity         int    `json:"quantity"`                     //【是】商品数量
	UnitPrice        int    `json:"unit_price"`                   //【是】商品单价
}

// PayTransactionsJsapiSceneInfo 场景信息
type PayTransactionsJsapiSceneInfo struct {
	PayerClientIp string                                  `json:"payer_client_ip"`      //【是】用户终端IP
	DeviceId      string                                  `json:"device_id,omitempty"`  //【否】商户端设备号
	StoreInfo     *PayTransactionsJsapiSceneInfoStoreInfo `json:"store_info,omitempty"` //【否】商户门店信息
}

// PayTransactionsJsapiSceneInfoStoreInfo 商户门店信息
type PayTransactionsJsapiSceneInfoStoreInfo struct {
	Id       string `json:"id"`                  //【是】门店编号
	Name     string `json:"name,omitempty"`      //【否】门店名称
	AreaCode string `json:"area_code,omitempty"` //【否】地区编码
	Address  string `json:"address,omitempty"`   //【否】详细地址
}

// PayTransactionsJsapiSettleInfo 结算信息
type PayTransactionsJsapiSettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` //【否】是否指定分账
}

// PayTransactionsJsapiResult 返回参数
type PayTransactionsJsapiResult struct {
	PrepayId string `json:"prepay_id"`
}

// PayTransactionsJsapi 小程序 JSAPI下单 https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
func (app *App) PayTransactionsJsapi(param PayTransactionsJsapi) (resp PayTransactionsJsapiResult, result ErrResp, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}

	body, result, err := app.request("pay/transactions/jsapi", params, "POST")

	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return
	}
	return
}
