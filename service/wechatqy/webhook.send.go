package wechatqy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

type WebhookSend struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// WebhookSend 群机器人 发送消息
// https://developer.work.weixin.qq.com/document/path/99110
func (c *Client) WebhookSend(ctx context.Context, key string, Type string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("cgi-bin/webhook/send?key=%s&type=%s", key, Type), params, http.MethodPost, &response)
	return
}

// WebhookSendURL 群机器人 发送消息
// https://developer.work.weixin.qq.com/document/path/99110
func (c *Client) WebhookSendURL(ctx context.Context, url string, Type string, notMustParams ...*gorequest.Params) (response WebhookSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, fmt.Sprintf("%s&type=%s", url, Type), params, http.MethodPost, &response)
	return
}

type MarkdownFormatDetails struct {
	Label string `json:"label"`           // 键
	Value string `json:"value"`           // 值
	Color string `json:"color,omitempty"` // 颜色，可选
}

func MarkdownFormat(ctx context.Context, title string, details []MarkdownFormatDetails) (markdownContent string) {

	// 使用 []string 动态存储每一行的内容
	var markdownLines []string

	// 添加标题内容
	markdownLines = append(markdownLines, title)

	// 动态添加详细信息
	for _, detail := range details {
		if detail.Color != "" {
			// 如果有颜色，则添加颜色标记
			line := fmt.Sprintf("> %s：<font color=\"%s\">%s</font>", detail.Label, detail.Color, detail.Value)
			markdownLines = append(markdownLines, line)
		} else {
			// 如果没有颜色，则直接拼接文本
			line := fmt.Sprintf("> %s：%s", detail.Label, detail.Value)
			markdownLines = append(markdownLines, line)
		}
	}

	// 将所有内容拼接成最终的 markdown 字符串
	markdownContent = strings.Join(markdownLines, "\n")

	return markdownContent
}
