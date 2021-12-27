package kashangwl

// ProductResult 返回参数
type ProductResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id                      int64   `json:"id"`
		ProductName             string  `json:"product_name,omitempty"`
		Name                    string  `json:"name"`
		Price                   float64 `json:"price"`
		ValidPurchasingQuantity string  `json:"valid_purchasing_quantity"`
		SuperiorCommissionsRate int     `json:"superior_commissions_rate"`
		Type                    int     `json:"type"`
		SupplyState             int     `json:"supply_state"`
		StockState              int     `json:"stock_state"`
		BanStartAt              string  `json:"ban_start_at,omitempty"`
		BanEndAt                string  `json:"ban_end_at,omitempty"`
	} `json:"data"`
}

// Product 获取单个商品信息
// http://doc.cqmeihu.cn/sales/product-info.html
func (app App) Product(productId int64) (body []byte, err error) {
	// 参数
	params := NewParams()
	params.Set("product_id", productId)
	// 请求
	body, err = app.request("http://www.kashangwl.com/api/product", params)
	return body, err
}

// ProductRechargeParamsResult 返回参数
type ProductRechargeParamsResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RechargeAccountLabel string `json:"recharge_account_label"`
		RechargeParams       []struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Options string `json:"options"`
		} `json:"recharge_params"`
	} `json:"data"`
}

// ProductRechargeParams 接口说明
// 获取商品的充值参数（仅支持充值类商品）
// http://doc.cqmeihu.cn/sales/ProductParams.html
func (app App) ProductRechargeParams(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("http://www.kashangwl.com/api/product/recharge-params", params)
	return body, err
}
