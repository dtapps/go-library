package meituan

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseServeHttpOrder 返回参数
type ResponseServeHttpOrder struct {
	Smstitle            string `json:"appId,omitempty"`               // 订单标题
	Quantity            string `json:"CreateTime,omitempty"`          // 订单数量
	Orderid             string `json:"orderid,omitempty"`             // 订单id
	Dealid              string `json:"dealid,omitempty"`              // 店铺id（部分存在）
	Paytime             string `json:"paytime,omitempty"`             // 订单支付时间，10位时间戳
	ActId               string `json:"actId,omitempty"`               // 活动id，可以在联盟活动列表中查看获取
	BusinessLine        string `json:"businessLine,omitempty"`        // 详见业务线类型
	SubBusinessLine     string `json:"subBusinessLine,omitempty"`     // 子业务线
	Type                string `json:"type,omitempty"`                // 订单类型，枚举值同订单查询接口定义
	Ordertime           string `json:"ordertime,omitempty"`           // 下单时间，10位时间戳
	Sid                 string `json:"sid,omitempty"`                 // 媒体推广位sid
	Appkey              string `json:"appkey,omitempty"`              // 媒体名称，可在推广者备案-媒体管理中查询
	Uid                 string `json:"uid,omitempty"`                 // 渠道id
	Status              string `json:"status,omitempty"`              // 订单状态，枚举值同订单查询接口返回定义
	Total               string `json:"total,omitempty"`               // 订单总金额
	PayPrice            string `json:"payPrice,omitempty"`            // 订单实际支付金额
	ModTime             string `json:"modTime,omitempty"`             // 订单修改时间
	ProductId           string `json:"productId,omitempty"`           // 商品ID
	ProductName         string `json:"productName,omitempty"`         // 商品名称
	Direct              string `json:"direct,omitempty"`              // 订单实际支付金额
	Ratio               string `json:"ratio,omitempty"`               // 订单返佣比例，cps活动的订单会返回该字段
	Sign                string `json:"sign,omitempty"`                // 订单签名字段，计算方法参见文档中签名(sign)生成逻辑
	TradeTypeList       string `json:"tradeTypeList,omitempty"`       // 优选订单类型返回该字段
	ConsumeType         string `json:"consumeType,omitempty"`         // 核销类型
	RefundType          string `json:"refundType,omitempty"`          // 退款类型
	EncryptionVoucherId string `json:"encryptionVoucherId,omitempty"` // 消费券加密券ID
}

// ServeHttpOrder 订单回推接口（新版）
// https://union.meituan.com/v2/apiDetail?id=22
func (app *App) ServeHttpOrder(c *gin.Context) (validateJson ResponseServeHttpOrder, err error) {
	// 声明接收的变量
	err = c.ShouldBind(&validateJson)
	return
}

// Success 数据正常
func (r *ResponseServeHttpOrder) Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "ok",
	})
}

// Error 数据错误
func (r *ResponseServeHttpOrder) Error(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 1,
		"errmsg":  "err",
	})
}
