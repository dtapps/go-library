package wikeyun

import (
	"encoding/json"
	"errors"
)

// RechargePpushOrderRequest 请求参数
type RechargePpushOrderRequest struct {
	StoreId     string `json:"store_id"`         // 铺ID
	Mobile      string `json:"mobile"`           // 充值号码
	OrderNo     string `json:"order_no"`         // 充值订单号
	Money       int    `json:"money"`            // 充值金额（100，200）
	RechargeTyp int    `json:"recharge_typ"`     // 是 1快充 0慢充
	NotifyUrl   string `json:"notify_url"`       // 异步回调地址（POST）
	Change      int    `json:"change,omitempty"` // 失败更换渠道充值 0 否 1是
	Source      int    `json:"source,omitempty"` // 是否强制渠道 传source字段则可以强制某渠道，强制快充走94折则，source传6
}

// RechargePpushOrderResponse 返回参数
type RechargePpushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

// RechargePushOrder 充值请求业务参数
func (app *App) RechargePushOrder(notMustParams ...Params) (result RechargePpushOrderResponse, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("https://router.wikeyun.cn/rest/Recharge/pushOrder", params)
	if err != nil {
		return result, errors.New(err.Error())
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.New(err.Error())
	}
	return result, nil
}
