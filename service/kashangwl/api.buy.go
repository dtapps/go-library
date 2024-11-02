package kashangwl

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiBuyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderId      string `json:"order_id"`      // 订单号
		ProductPrice string `json:"product_price"` // 商品价格
		TotalPrice   string `json:"total_price"`   // 总支付价格
		RechargeUrl  string `json:"recharge_url"`  // 卡密充值网址
		State        int    `json:"state"`         // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
		Cards        []struct {
			CardNo       string `json:"card_no"`
			CardPassword string `json:"card_password"`
		} `json:"cards,omitempty"` // 卡密（仅当订单成功并且商品类型为卡密时返回此数据）
		Tickets []struct {
			No     string `json:"no"`
			Ticket string `json:"ticket"`
		} `json:"tickets,omitempty"` // 卡券（仅当订单成功并且商品类型为卡券时返回此数据）
	} `json:"data"`
}

type ApiBuyResult struct {
	Result ApiBuyResponse     // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newApiBuyResult(result ApiBuyResponse, body []byte, http gorequest.Response) *ApiBuyResult {
	return &ApiBuyResult{Result: result, Body: body, Http: http}
}

// ApiBuy 购买商品
// product_id = 商品编号
// recharge_account = 充值账号
// recharge_template_input_items = 模板充值参数
// quantity = 购买数量
// notify_url = 异步通知地址
// outer_order_id = 外部订单号
// safe_cost = 安全进价
// client_ip = 购买的用户真实IP
// http://doc.cqmeihu.cn/sales/buy.html
func (c *Client) ApiBuy(ctx context.Context, productID int64, quantity int64, notMustParams ...gorequest.Params) (*ApiBuyResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("product_id", productID) // 商品编号
	params.Set("quantity", quantity)    // 购买数量

	// 请求
	var response ApiBuyResponse
	request, err := c.request(ctx, "api/buy", params, &response)
	return newApiBuyResult(response, request.ResponseBody, request), err
}
