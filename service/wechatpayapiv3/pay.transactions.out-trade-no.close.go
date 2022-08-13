package wechatpayapiv3

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayTransactionsOutTradeNoCloseResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
	Err  error              // 错误
}

func newPayTransactionsOutTradeNoCloseResult(body []byte, http gorequest.Response, err error) *PayTransactionsOutTradeNoCloseResult {
	return &PayTransactionsOutTradeNoCloseResult{Body: body, Http: http, Err: err}
}

// PayTransactionsOutTradeNoClose 关闭订单API https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (c *Client) PayTransactionsOutTradeNoClose(ctx context.Context, OutTradeNo string) *PayTransactionsOutTradeNoCloseResult {
	// 参数
	params := gorequest.NewParams()
	params["mchid"] = c.GetMchId()
	// 	请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/v3/pay/transactions/out-trade-no/%s/close", OutTradeNo), params, http.MethodPost, false)
	return newPayTransactionsOutTradeNoCloseResult(request.ResponseBody, request, err)
}
