package kashangwl

import "encoding/json"

type ApiProductResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id                      int     `json:"id"`
		ProductName             string  `json:"product_name"`
		Name                    string  `json:"name"`
		Price                   float64 `json:"price"`
		ValidPurchasingQuantity string  `json:"valid_purchasing_quantity"`
		SuperiorCommissionsRate int     `json:"superior_commissions_rate"`
		Type                    int     `json:"type"`
		SupplyState             int     `json:"supply_state"`
		StockState              int     `json:"stock_state"`
		BanStartAt              string  `json:"ban_start_at"`
		BanEndAt                string  `json:"ban_end_at"`
	} `json:"data"`
}

type ApiProductResult struct {
	Result ApiProductResponse // 结果
	Body   []byte             // 内容
	Err    error              // 错误
}

func NewApiProductResult(result ApiProductResponse, body []byte, err error) *ApiProductResult {
	return &ApiProductResult{Result: result, Body: body, Err: err}
}

// ApiProduct 获取单个商品信息
// http://doc.cqmeihu.cn/sales/product-info.html
func (app App) ApiProduct(productId int64) *ApiProductResult {
	// 参数
	params := NewParams()
	params.Set("product_id", productId)
	// 请求
	body, err := app.request("http://www.kashangwl.com/api/product", params)
	// 定义
	var response ApiProductResponse
	err = json.Unmarshal(body, &response)
	return NewApiProductResult(response, body, err)
}