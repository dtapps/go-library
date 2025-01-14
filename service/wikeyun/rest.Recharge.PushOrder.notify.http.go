package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"net/http"
)

// ResponseRestRechargePushOrderNotifyHttp 声明接收的变量
type ResponseRestRechargePushOrderNotifyHttp struct {
	Status      int64   `json:"status,omitempty"`       // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	Mobile      string  `json:"mobile,omitempty"`       // 充值手机号
	OrderNo     string  `json:"order_no,omitempty"`     // 第三方单号
	OrderNumber string  `json:"order_number,omitempty"` // 微客云平台单号
	Amount      string  `json:"amount,omitempty"`       // 充值金额，如50，100，200可选
	Fanli       float64 `json:"fanli,omitempty"`        // 返利金额
	CostPrice   float64 `json:"cost_price,omitempty"`   // 成本价格
	Sign        string  `json:"sign,omitempty"`         // 加密内容
}

// RestRechargePushOrderNotifyHttp 话费充值推送 - 回调通知
// https://open.wikeyun.cn/#/document/1/article/302
func (c *Client) RestRechargePushOrderNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateJson ResponseRestRechargePushOrderNotifyHttp, err error) {
	err = gojson.NewDecoder(r.Body).Decode(&validateJson)
	return validateJson, err
}
