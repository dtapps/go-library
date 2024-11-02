package wechatpayapiv3

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayTransactionsOutTradeNoCloseResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
}

func newPayTransactionsOutTradeNoCloseResult(body []byte, http gorequest.Response) *PayTransactionsOutTradeNoCloseResult {
	return &PayTransactionsOutTradeNoCloseResult{Body: body, Http: http}
}

// PayTransactionsOutTradeNoClose 关闭订单API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (c *Client) PayTransactionsOutTradeNoClose(ctx context.Context, OutTradeNo string, notMustParams ...gorequest.Params) (*PayTransactionsOutTradeNoCloseResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mchid", c.GetMchId())

	// 	请求
	request, err := c.request(ctx, fmt.Sprintf("v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost, false, nil)
	return newPayTransactionsOutTradeNoCloseResult(request.ResponseBody, request), err
}
