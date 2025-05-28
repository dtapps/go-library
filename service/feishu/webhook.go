package feishu

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"go.dtapp.net/library/utils/gorequest"
)

type WebhookSend struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// WebhookSend 发送消息
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSend(ctx context.Context, key string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params, &response)
	return
}

// WebhookSendURL 发送消息
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendURL(ctx context.Context, url string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, url, params, &response)
	return
}

// WebhookSendSign 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSign(ctx context.Context, key string, secret string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("timestamp", time.Now().Unix())
	sign, _ := c.webhookSendSignGenSign(secret, fmt.Sprintf("%v", params.Get("timestamp")))
	params.Set("sign", sign)

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params, &response)
	return
}

// WebhookSendSignURL 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSignURL(ctx context.Context, url string, secret string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("timestamp", time.Now().Unix())
	sign, _ := c.webhookSendSignGenSign(secret, fmt.Sprintf("%v", params.Get("timestamp")))
	params.Set("sign", sign)

	// 请求
	err = c.request(ctx, url, params, &response)
	return
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

type MarkdownFormatDetails struct {
	Label string `json:"label"` // 键
	Value string `json:"value"` // 值
}

type MarkdownFormatResponse struct {
	Title   string                            `json:"title"`
	Content [][]MarkdownFormatResponseContent `json:"content"`
}
type MarkdownFormatResponseContent struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
}

func MarkdownFormat(ctx context.Context, title string, details []MarkdownFormatDetails) (response MarkdownFormatResponse) {

	// 添加标题内容
	response.Title = title

	// 动态添加详细信息
	content := make([]MarkdownFormatResponseContent, 0)
	for i, detail := range details {
		if i < len(details)-1 {
			// 非最后一行，添加换行符
			content = append(content, MarkdownFormatResponseContent{
				Tag:  "text",
				Text: fmt.Sprintf(" %s：%s\n", detail.Label, detail.Value),
			})
		} else {
			// 最后一行，不添加换行符
			content = append(content, MarkdownFormatResponseContent{
				Tag:  "text",
				Text: fmt.Sprintf(" %s：%s", detail.Label, detail.Value),
			})
		}
	}
	response.Content = append(response.Content, content)

	return response
}
