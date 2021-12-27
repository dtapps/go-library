package wechatqy

import (
	"encoding/json"
	"fmt"
)

// WebhookSendText 文本类型
type WebhookSendText struct {
	Msgtype string `json:"msgtype"` // 消息类型，此时固定为text
	Text    struct {
		Content             string   `json:"content"`               // 文本内容，最长不超过2048个字节，必须是utf8编码
		MentionedList       []string `json:"mentioned_list"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
		MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
	} `json:"text"`
}

type WebhookSendResult struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// WebhookSend https://work.weixin.qq.com/api/doc/90000/90136/91770
func (app *App) WebhookSend(notMustParams ...Params) (result WebhookSendResult, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s&type=%s", app.Key, "text"), params)
	if err != nil {
		return WebhookSendResult{}, err
	}
	err = json.Unmarshal(request, &result)
	if err != nil {
		return WebhookSendResult{}, err
	}
	return result, err
}
