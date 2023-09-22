package kashangwl

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ApiOuterOrderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id                 string  `json:"id"`                   // 订单号
		ProductId          int     `json:"product_id"`           // 商品编号
		ProductName        string  `json:"product_name"`         // 商品名称
		ProductType        int     `json:"product_type"`         // 商品类型（1：充值，2：卡密，3：卡券，4：人工）
		ProductPrice       string  `json:"product_price"`        // 售价
		Quantity           int     `json:"quantity"`             // 购买数量
		TotalPrice         string  `json:"total_price"`          // 总支付价格
		RefundedAmount     float64 `json:"refunded_amount"`      // 已退款金额
		BuyerCustomerId    int     `json:"buyer_customer_id"`    // 买家编号
		BuyerCustomerName  string  `json:"buyer_customer_name"`  // 买家名称
		SellerCustomerId   int     `json:"seller_customer_id"`   // 卖家编号
		SellerCustomerName string  `json:"seller_customer_name"` // 卖家名称
		State              int     `json:"state"`                // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
		CreatedAt          string  `json:"created_at"`           // 下单时间
		RechargeAccount    string  `json:"recharge_account"`     // 充值账号
		ProgressInit       int     `json:"progress_init"`        // 充值进度：初始值
		ProgressNow        int     `json:"progress_now"`         // 充值进度：现在值
		ProgressTarget     int     `json:"progress_target"`      // 充值进度：目标值
		RechargeInfo       string  `json:"recharge_info"`        // 返回信息
		RechargeUrl        string  `json:"recharge_url"`         // 卡密充值网址
		Cards              []struct {
			No       string `json:"no"`
			Password string `json:"password"`
		} `json:"cards"` //【卡密类订单】卡密
		RechargeParams string `json:"recharge_params"`          //【充值类订单】
		OuterOrderId   string `json:"outer_order_id,omitempty"` // 外部订单号
	} `json:"data"`
}

type ApiOuterOrderResult struct {
	Result ApiOuterOrderResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newApiOuterOrderResult(result ApiOuterOrderResponse, body []byte, http gorequest.Response) *ApiOuterOrderResult {
	return &ApiOuterOrderResult{Result: result, Body: body, Http: http}
}

// ApiOuterOrder 使用外部订单号获取单个订单信息
// 仅能获取自己购买的订单
// http://doc.cqmeihu.cn/sales/outer-order-info.html
func (c *Client) ApiOuterOrder(ctx context.Context, notMustParams ...*gorequest.Params) (*ApiOuterOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/outer-order", params)
	if err != nil {
		return newApiOuterOrderResult(ApiOuterOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiOuterOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiOuterOrderResult(response, request.ResponseBody, request), err
}

func (resp ApiOuterOrderResponse) GetStateDesc(state int) string {
	switch state {
	case 100:
		return "等待发货"
	case 101:
		return "正在充值"
	case 200:
		return "交易成功"
	case 500:
		return "交易失败"
	case 501:
		return "未知状态"
	}
	return ""
}
