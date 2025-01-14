package feishu

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
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
	Http   gorequest.Response  // 请求
}

func newWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest.Response) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http}
}

// WebhookSend 发送消息
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSend(ctx context.Context, key string, notMustParams ...*gorequest.Params) (*WebhookSendResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

// WebhookSendURL 发送消息
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendURL(ctx context.Context, url string, notMustParams ...*gorequest.Params) (*WebhookSendResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, url, params, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

// WebhookSendSign 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSign(ctx context.Context, key string, secret string, notMustParams ...*gorequest.Params) (*WebhookSendResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("timestamp", gotime.Current().Timestamp())
	sign, _ := c.webhookSendSignGenSign(secret, fmt.Sprintf("%v", params.Get("timestamp")))
	params.Set("sign", sign)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

// WebhookSendSignURL 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSignURL(ctx context.Context, url string, secret string, notMustParams ...*gorequest.Params) (*WebhookSendResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("timestamp", gotime.Current().Timestamp())
	sign, _ := c.webhookSendSignGenSign(secret, fmt.Sprintf("%v", params.Get("timestamp")))
	params.Set("sign", sign)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, url, params, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

func (c *Client) webhookSendSignGenSign(secret string, timestamp string) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
