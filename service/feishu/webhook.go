package feishu

import (
	"encoding/json"
	"fmt"
)

type WebhookSendResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type WebhookSendResult struct {
	Result WebhookSendResponse // 结果
	Body   []byte              // 内容
	Err    error               // 错误
}

func NewWebhookSendResult(result WebhookSendResponse, body []byte, err error) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Err: err}
}

// WebhookSend https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (app *App) WebhookSend(notMustParams ...Params) *WebhookSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://open.feishu.cn/open-apis/bot/v2/hook/%s", app.Key), params)
	// 定义
	var response WebhookSendResponse
	err = json.Unmarshal(body, &response)
	return NewWebhookSendResult(response, body, err)
}
