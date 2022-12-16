package wikeyun

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ResponseRestPowerPushOrderNotifyGin struct {
	Status        int     `form:"status" json:"status" xml:"status" uri:"status" binding:"omitempty"`                                // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
	ArrivedAmount string  `form:"arrived_amount" json:"arrived_amount" xml:"arrived_amount" uri:"arrived_amount" binding:"required"` // 到账金额
	OrderNo       string  `form:"order_no" json:"order_no" xml:"order_no" uri:"order_no" binding:"required"`                         // 第三方单号
	OrderNumber   string  `form:"order_number" json:"order_number" xml:"order_number" uri:"order_number" binding:"required"`         // 微客云平台单号
	Amount        string  `form:"amount" json:"amount" xml:"amount" uri:"amount" binding:"required"`                                 // 充值金额，如50，100，200可选
	Fanli         float64 `form:"fanli" json:"fanli" xml:"fanli" uri:"fanli" binding:"required"`                                     // 返利金额
	CostPrice     float64 `form:"cost_price" json:"cost_price" xml:"cost_price" uri:"cost_price" binding:"required"`                 // 成本价格
	Sign          string  `form:"sign" json:"sign" xml:"sign" uri:"sign" binding:"omitempty"`                                        // 加密内容
	FailReason    string  `form:"failReason" json:"failReason" xml:"failReason" uri:"failReason" binding:"omitempty"`                // 失败原因，有些渠道没有返回，不是很准确，但电费失败大部分原因都是户号不对或者地区不对或者缴费金额小于欠费金额
}

// RestPowerPushOrderNotifyGin 电费充值API - 回调通知
// https://open.wikeyun.cn/#/document/1/article/303
func (c *Client) RestPowerPushOrderNotifyGin(ctx context.Context, ginCtx *gin.Context) (ResponseRestPowerPushOrderNotifyGin, error) {

	// 声明接收的变量
	var validateJson struct {
		Status        int     `form:"status" json:"status" xml:"status" uri:"status" binding:"omitempty"`                                // 状态 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8 已存单 9 已取消 10返销 11部分到账
		ArrivedAmount string  `form:"arrived_amount" json:"arrived_amount" xml:"arrived_amount" uri:"arrived_amount" binding:"required"` // 到账金额
		OrderNo       string  `form:"order_no" json:"order_no" xml:"order_no" uri:"order_no" binding:"required"`                         // 第三方单号
		OrderNumber   string  `form:"order_number" json:"order_number" xml:"order_number" uri:"order_number" binding:"required"`         // 微客云平台单号
		Amount        string  `form:"amount" json:"amount" xml:"amount" uri:"amount" binding:"required"`                                 // 充值金额，如50，100，200可选
		Fanli         float64 `form:"fanli" json:"fanli" xml:"fanli" uri:"fanli" binding:"required"`                                     // 返利金额
		CostPrice     float64 `form:"cost_price" json:"cost_price" xml:"cost_price" uri:"cost_price" binding:"required"`                 // 成本价格
		Sign          string  `form:"sign" json:"sign" xml:"sign" uri:"sign" binding:"omitempty"`                                        // 加密内容
		FailReason    string  `form:"failReason" json:"failReason" xml:"failReason" uri:"failReason" binding:"omitempty"`                // 失败原因，有些渠道没有返回，不是很准确，但电费失败大部分原因都是户号不对或者地区不对或者缴费金额小于欠费金额
	}

	err := ginCtx.ShouldBind(&validateJson)

	return validateJson, err
}
