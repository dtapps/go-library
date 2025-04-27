package kashangwl

import (
	"context"
	"encoding/json"
	"net/http"
)

// ResponseApiBuyNotifyHttp 购买商品 - 回调通知 - 请求参数
type ResponseApiBuyNotifyHttp struct {
	OrderId      string `json:"order_id"`                 // 订单编号
	OuterOrderId string `json:"outer_order_id,omitempty"` // 商户订单号
	ProductId    int64  `json:"product_id,omitempty"`     // 商品编号
	Quantity     int64  `json:"quantity"`                 // 购买数量
	State        int64  `json:"state"`                    // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
	StateInfo    string `json:"state_info,omitempty"`     // 状态信息
	CreatedAt    string `json:"created_at"`               // 购买数量
}

// ApiBuyNotifyHttp 购买商品 - 回调通知
// http://doc.cqmeihu.cn/sales/order-status-notify.html
func (c *Client) ApiBuyNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateJson ResponseApiBuyNotifyHttp, err error) {
	err = json.NewDecoder(r.Body).Decode(&validateJson)
	return validateJson, err
}
