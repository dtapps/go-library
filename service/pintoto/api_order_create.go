package pintoto

import "encoding/json"

type ApiOrderCreate struct {
	ShowId           string `json:"showId"`                  // 排期的showId,由影院接口得来
	Seat             string `json:"seat"`                    // 用户所选的座位，例：1排1座,1排2座 以英文的逗号 “ , “隔开。 如果座位是情侣座，请传入 ： 1排1座(情侣座),1排2座(情侣座)
	ReservedPhone    string `json:"reservedPhone,omitempty"` // 下单时预留的手机号，方便问题沟通
	ThirdOrderId     string `json:"thirdOrderId"`            // 接入方的订单号， 接入方须保证此订单号唯一性
	NotifyUrl        string `json:"notifyUrl"`               // 回调地址，各个场景发生时，将通过此地址通知接入方，详情请看【回调api】
	AcceptChangeSeat int    `json:"acceptChangeSeat"`        // 是否允许调座，1-允许，0-不允许
	SeatId           string `json:"seatId,omitempty"`        // 座位接口的seatId字段， 如果有多个，则以竖线分割
	SeatNo           string `json:"seatNo,omitempty"`        // 座位接口的seatNo字段，如果有多个，则以竖线分割
}

type ApiOrderCreateResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// ApiOrderCreate 下单api https://www.showdoc.com.cn/1154868044931571/5891022916496848
func (app *App) ApiOrderCreate(param ApiOrderCreate) (result ApiOrderCreateResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("https://movieapi2.pintoto.cn/api/order/create", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
