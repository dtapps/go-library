package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type RechargeResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
}

type Recharge struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   struct {
		CreateTime  int64  `json:"create_time,omitempty"` // 下单时间
		Guishu      string `json:"guishu,omitempty"`      // 归属地
		Id          int64  `json:"id,omitempty"`          // 编号
		Mobile      string `json:"mobile"`                // 充值手机号
		OrderNumber string `json:"order_number"`          // 系统定单号
		OutTradeNum string `json:"out_trade_num"`         // 商户订单号
		ProductId   int64  `json:"product_id"`            // 产品ID
		Title       string `json:"title"`                 // 充值产品说明
		TotalPrice  string `json:"total_price"`           // 消费金额
	} `json:"data,omitempty"`
}

// Recharge 充值提交接口
// out_trade_num = 商户订单号，由商户自己生成唯一单号。（同一商户，不能存在相同单号订单，相同订单号不能提单）
// product_id = 产品ID（代理后台查看）
// mobile = 充值号码（手机号、电费户、qq号等）
// notify_url = 回调地址，用于接收充值状态回调
// amount = 面值，（不传不校验）如果产品的面值与此参数不同，提单驳回
// price = 最高成本，（不传不校验）如果产品成本超过这个值，提单驳回
// area = 电费省份/直辖市，如：四川、北京、上海，仅电费带此参数
// ytype = 电费验证三要素，1-身份证后6位，2-银行卡后六位,3-营业执照后六位，仅南网电费带此参数
// id_card_no = 身份证后6位/银行卡后6位/营业执照后6位，仅南网电费带此参数
// city = 地级市名，仅部分南网电费带此参数，是否带此参数需咨询渠道方
// param1 = 扩展参数，后台查看提交的产品类目是否需要提交此参数
// param2 = 扩展参数，后台查看提交的产品类目是否需要提交此参数
// param3 = 扩展参数，后台查看提交的产品类目是否需要提交此参数
// https://www.showdoc.com.cn/dyr/9227003154511692
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097250
func (c *Client) Recharge(ctx context.Context, outTradeNum string, productID int64, mobile string, notifyUrl string, notMustParams ...*gorequest.Params) (response Recharge, apiErr ErrorResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("out_trade_num", outTradeNum) // 商户订单号
	params.Set("product_id", productID)      // 产品ID
	params.Set("mobile", mobile)             // 充值号码
	params.Set("notify_url", notifyUrl)      // 回调地址
	params.Set("userid", c.GetUserID())      // 商户ID

	// 请求
	err = c.requestAndErr(ctx, "index/recharge", params, &response, &apiErr)
	return
}
