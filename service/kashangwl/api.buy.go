package kashangwl

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiBuyResponse struct {
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

type ApiBuyResult struct {
	Result ApiBuyResponse     // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newApiBuyResult(result ApiBuyResponse, body []byte, http gorequest.Response, err error) *ApiBuyResult {
	return &ApiBuyResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiBuy 购买商品
// http://doc.cqmeihu.cn/sales/buy.html
func (c *Client) ApiBuy(notMustParams ...gorequest.Params) *ApiBuyResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/api/buy", params)
	// 定义
	var response ApiBuyResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiBuyResult(response, request.ResponseBody, request, err)
}
