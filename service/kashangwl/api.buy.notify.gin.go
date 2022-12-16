package kashangwl

import (
	"context"
	"github.com/gin-gonic/gin"
)

// ResponseApiBuyNotifyGin 购买商品 - 回调通知 - 请求参数
type ResponseApiBuyNotifyGin struct {
	OrderId      string `form:"order_id" json:"order_id" xml:"order_id" uri:"order_id" binding:"required"`                          // 订单编号
	OuterOrderId string `form:"outer_order_id" json:"outer_order_id" xml:"outer_order_id" uri:"outer_order_id" binding:"omitempty"` // 商户订单号
	ProductId    int    `form:"product_id" json:"product_id" xml:"product_id" uri:"product_id" binding:"omitempty"`                 // 商品编号
	Quantity     int    `form:"quantity" json:"quantity" xml:"quantity" uri:"quantity" binding:"required"`                          // 购买数量
	State        int    `form:"state" json:"state" xml:"state" uri:"state" binding:"required"`                                      // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
	StateInfo    string `form:"state_info" json:"state_info" xml:"state_info" uri:"state_info" binding:"omitempty"`                 // 状态信息
	CreatedAt    string `form:"created_at" json:"created_at" xml:"created_at" uri:"created_at" binding:"required"`                  // 购买数量
}

// ApiBuyNotifyGin 购买商品 - 回调通知
// http://doc.cqmeihu.cn/sales/order-status-notify.html
func (c *Client) ApiBuyNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson ResponseApiBuyNotifyGin, err error) {

	err = ginCtx.ShouldBind(&validateJson)

	return validateJson, err
}
