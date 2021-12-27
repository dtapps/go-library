package kashangwl

// OrderResult 返回参数
type OrderResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id                int64  `json:"id"`
		ProductId         int    `json:"product_id"`
		ProductName       string `json:"product_name"`
		ProductType       int    `json:"product_type"`
		ProductPrice      string `json:"product_price"`
		Quantity          int    `json:"quantity"`
		TotalPrice        string `json:"total_price"`
		RefundedAmount    string `json:"refunded_amount"`
		BuyerCustomerId   int    `json:"buyer_customer_id"`
		BuyerCustomerName string `json:"buyer_customer_name"`
		State             int    `json:"state"`
		CreatedAt         string `json:"created_at"`
		OuterOrderId      string `json:"outer_order_id"`
		RechargeAccount   string `json:"recharge_account"`
		RechargeParams    string `json:"recharge_params"`
		RechargeInfo      string `json:"recharge_info"`
		RechargeUrl       string `json:"recharge_url"`
		Cards             []struct {
			No       string `json:"no"`
			Password string `json:"password"`
		} `json:"cards"`
	} `json:"data"`
}

// Order 获取单个订单信息。
// 仅能获取自己购买的订单。
// http://doc.cqmeihu.cn/sales/OrderInfo.html
func (app App) Order(orderId string) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("order_id", orderId)
	params := app.NewParamsWith(param)
	// 请求
	body, err = app.request("http://www.kashangwl.com/api/order", params)
	return body, err
}
