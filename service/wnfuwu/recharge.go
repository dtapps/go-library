package wnfuwu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RechargeResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   struct {
		OrderNumber string `json:"order_number"`  // 系统定单号
		Mobile      string `json:"mobile"`        // 充值手机号
		ProductId   string `json:"product_id"`    // 产品ID
		TotalPrice  string `json:"total_price"`   // 消费金额
		OutTradeNum string `json:"out_trade_num"` // 商户订单号
		Title       string `json:"title"`         // 充值产品说明
	} `json:"data"`
}

type RechargeResult struct {
	Result RechargeResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRechargeResult(result RechargeResponse, body []byte, http gorequest.Response) *RechargeResult {
	return &RechargeResult{Result: result, Body: body, Http: http}
}

// Recharge 充值提交接口
// https://www.showdoc.com.cn/dyr/9227003154511692
func (c *Client) Recharge(ctx context.Context, notMustParams ...gorequest.Params) (*RechargeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId())
	// 请求
	request, err := c.request(ctx, apiUrl+"/index/recharge", params)
	if err != nil {
		return newRechargeResult(RechargeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RechargeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRechargeResult(response, request.ResponseBody, request), err
}
