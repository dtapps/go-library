package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type CheckResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
}

type Check struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   []struct {
		Id           int64  `json:"id,omitempty"`           // 编号
		OrderNumber  string `json:"order_number"`           // 系统定单号
		Status       int64  `json:"status"`                 // 充值状态：-1取消，0充值中 ，1充值成功，2充值失败，3部分成功
		OutTradeNum  string `json:"out_trade_num"`          // 商户订单号
		CreateTime   int64  `json:"create_time"`            // 下单时间
		Mobile       string `json:"mobile"`                 // 手机号
		ProductId    int64  `json:"product_id"`             // 产品ID
		ChargeAmount string `json:"charge_amount"`          // 充值成功面额
		ChargeKami   string `json:"charge_kami"`            // 卡密流水
		Isp          string `json:"isp,omitempty"`          // 运营商
		ProductName  string `json:"product_name,omitempty"` // 产品名称
		FinishTime   int64  `json:"finish_time,omitempty"`  // 完成时间
		Remark       string `json:"remark,omitempty"`       // 备注
		State        int64  `json:"state"`                  // 充值状态：-1取消，0充值中 ，1充值成功，2充值失败，3部分成功
		Voucher      string `json:"voucher,omitempty"`      // 凭证
	} `json:"data,omitempty"`
}

// Check 自发查询订单状态
// out_trade_nums = 商户订单号；多个用英文,分割
// https://www.showdoc.com.cn/dyr/9227006175502841
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097254
func (c *Client) Check(ctx context.Context, outTradeNums string, notMustParams ...*gorequest.Params) (response Check, apiErr ErrorResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID())        // 账户ID
	params.Set("out_trade_nums", outTradeNums) // 商户订单号；多个用英文,分割

	// 请求
	err = c.requestAndErr(ctx, "index/check", params, &response, &apiErr)
	return
}
