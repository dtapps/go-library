package message

type Product struct {
	ProductId int `json:"product_id"` //true 商品编号
}

// ProductRechargeParams 获取商品的充值参数。
type ProductRechargeParams struct {
	ProductId int `json:"product_id"` //true 商品编号
}

// Buy 购买商品（不支持购买选号类型的商品）
type Buy struct {
	ProductId                  int         `json:"product_id"`                    //true 商品编号
	RechargeAccount            string      `json:"recharge_account"`              //false	充值账号
	RechargeTemplateInputItems interface{} `json:"recharge_template_input_items"` //false	模板充值参数（[1]见下方说明）
	Quantity                   int         `json:"quantity"`                      //true	购买数量
	NotifyUrl                  string      `json:"notify_url"`                    //false	异步通知地址
	OuterOrderId               string      `json:"outer_order_id"`                //false	外部订单号
}

// Order 获取单个订单信息
type Order struct {
	OrderId int64 `json:"order_id"` //true	订单号
}

// Customer 获取商家信息
type Customer struct {
	CustomerId int `json:"customer_id"` //true	商家编号
}
