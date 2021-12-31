package kashangwl

import "encoding/json"

type ProductResponse struct {
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

type ProductResult struct {
	Result ProductResponse // 结果
	Body   []byte          // 内容
	Err    error           // 错误
}

func NewProductResult(result ProductResponse, body []byte, err error) *ProductResult {
	return &ProductResult{Result: result, Body: body, Err: err}
}

// Product 获取单个商品信息
// http://doc.cqmeihu.cn/sales/product-info.html
func (app App) Product(productId int64) *ProductResult {
	// 参数
	params := NewParams()
	params.Set("product_id", productId)
	// 请求
	body, err := app.request("http://www.kashangwl.com/api/product", params)
	// 定义
	var response ProductResponse
	err = json.Unmarshal(body, &response)
	return NewProductResult(response, body, err)
}
