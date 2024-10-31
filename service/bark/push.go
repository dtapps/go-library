package bark

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PushResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type PushResult struct {
	Result PushResponse       // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newPushResult(result PushResponse, body []byte, http gorequest.Response) *PushResult {
	return &PushResult{Result: result, Body: body, Http: http}
}

// Push 推送消息
// https://bark.day.app/#/tutorial?id=json-%e8%af%b7%e6%b1%82-key-%e5%8f%af%e4%bb%a5%e6%94%be%e8%bf%9b%e8%af%b7%e6%b1%82%e4%bd%93%e4%b8%adurl-%e8%b7%af%e5%be%84%e9%a1%bb%e4%b8%ba-push%ef%bc%8c%e4%be%8b%e5%a6%82
func (c *Client) Push(ctx context.Context, notMustParams ...gorequest.Params) (*PushResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("device_key", c.config.pushKey)

	// 请求
	var response PushResponse
	request, err := c.request(ctx, apiUrl+"push", params, http.MethodPost, &response)
	return newPushResult(response, request.ResponseBody, request), err
}
