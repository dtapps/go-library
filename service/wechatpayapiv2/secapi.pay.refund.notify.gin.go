package wechatpayapiv2

import (
	"context"
	"github.com/gin-gonic/gin"
)

// SecApiPayRefundNotifyGinRequest 小程序支付 - 申请退款 - 回调通知 - 请求参数
type SecApiPayRefundNotifyGinRequest struct {
	ReturnCode string `form:"return_code" json:"return_code" xml:"return_code" uri:"return_code" binding:"required"` // 返回状态码
	ReturnMsg  string `form:"return_msg" json:"return_msg" xml:"return_msg" uri:"return_msg" binding:"omitempty"`    // 返回信息
	Appid      string `form:"appid" json:"appid" xml:"appid" uri:"appid" binding:"omitempty"`                        // 公众账号ID
	MchId      string `form:"mch_id" json:"mch_id" xml:"mch_id" uri:"mch_id" binding:"omitempty"`                    // 退款的商户号
	NonceStr   string `form:"nonce_str" json:"nonce_str" xml:"nonce_str" uri:"nonce_str" binding:"omitempty"`        // 随机字符串
	ReqInfo    string `form:"req_info" json:"req_info" xml:"req_info" uri:"req_info" binding:"omitempty"`            // 加密信息
}

// SecApiPayRefundNotifyGin 小程序支付 - 申请退款 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_16&index=10
func (c *Client) SecApiPayRefundNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson SecApiPayRefundNotifyGinRequest, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	return validateJson, err
}
