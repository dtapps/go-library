package kashangwl

import "encoding/json"

type BuyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderID      int64  `json:"order_id"`      // 订单号
		ProductPrice string `json:"product_price"` // 商品价格
		TotalPrice   string `json:"total_price"`   // 总支付价格
		RechargeUrl  string `json:"recharge_url"`  // 卡密充值网址
		State        int    `json:"state"`         // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
		Cards        []struct {
			CardNo       string `json:"card_no"`
			CardPassword string `json:"card_password"`
		} `json:"cards,omitempty"` // 	卡密（仅当订单成功并且商品类型为卡密时返回此数据）
		Tickets []struct {
			No     string `json:"no"`
			Ticket string `json:"ticket"`
		} `json:"tickets,omitempty"` // 	卡券（仅当订单成功并且商品类型为卡券时返回此数据）
	} `json:"data"`
}

type BuyResult struct {
	Result BuyResponse // 结果
	Body   []byte      // 内容
	Err    error       // 错误
}

func NewBuyResult(result BuyResponse, body []byte, err error) *BuyResult {
	return &BuyResult{Result: result, Body: body, Err: err}
}

// Buy 购买商品
// http://doc.cqmeihu.cn/sales/BuyProduct.html
func (app *App) Buy(notMustParams ...Params) *BuyResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("http://www.kashangwl.com/api/buy", params)
	// 定义
	var response BuyResponse
	err = json.Unmarshal(body, &response)
	return NewBuyResult(response, body, err)
}
