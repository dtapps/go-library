package dingtalk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.dtapp.net/library/utils/gorequest"
)

type RobotSend struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// RobotSend 发送消息
// https://open.dingtalk.com/document/dingstart/custom-bot-send-message-type
func (c *Client) RobotSend(ctx context.Context, accessToken string, notMustParams ...*gorequest.Params) (response RobotSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("/robot/send?access_token=%s", accessToken), params, http.MethodPost, &response)
	return
}

// RobotSendURL 发送消息
// https://open.dingtalk.com/document/dingstart/custom-bot-send-message-type
func (c *Client) RobotSendURL(ctx context.Context, url string, notMustParams ...*gorequest.Params) (response RobotSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, url, params, http.MethodPost, &response)
	return
}

// RobotSendSign 发送消息签名版
// https://open.dingtalk.com/document/dingstart/custom-bot-send-message-type
func (c *Client) RobotSendSign(ctx context.Context, accessToken string, secret string, notMustParams ...*gorequest.Params) (response RobotSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6

	// 请求
	err = c.request(ctx, apiUrl+fmt.Sprintf("robot/send?access_token=%s&timestamp=%d&sign=%s", accessToken, timestamp, c.robotSendSignGetSign(secret, timestamp)), params, http.MethodPost, &response)
	return
}

// RobotSendSignURL 发送消息签名版
// https://open.dingtalk.com/document/dingstart/custom-bot-send-message-type
func (c *Client) RobotSendSignURL(ctx context.Context, url string, secret string, notMustParams ...*gorequest.Params) (response RobotSend, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6

	// 请求
	err = c.request(ctx, fmt.Sprintf("%s&timestamp=%d&sign=%s", url, timestamp, c.robotSendSignGetSign(secret, timestamp)), params, http.MethodPost, &response)
	return
}

func (c *Client) robotSendSignGetSign(secret string, timestamp int64) string {
	secStr := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}

type MarkdownFormatDetails struct {
	Label string `json:"label"` // 键
	Value string `json:"value"` // 值
}

type MarkdownFormatResponse struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func MarkdownFormat(ctx context.Context, title string, details []MarkdownFormatDetails) (response MarkdownFormatResponse) {

	// 使用 []string 动态存储每一行的内容
	var markdownLines []string

	// 添加标题内容
	response.Title = title

	// 动态添加详细信息
	for _, detail := range details {
		line := fmt.Sprintf("> %s：%s", detail.Label, detail.Value)
		markdownLines = append(markdownLines, line)
	}
	// 将所有内容拼接成最终的 markdown 字符串
	response.Text = strings.Join(markdownLines, "\n")

	return response
}
