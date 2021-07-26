package qywechat

import (
	v20210726 "github.com/dtapps/go-library/qywechat/v20210726"
	"gopkg.in/dtapps/go-library.v2/qywechat/message"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bot := v20210726.QyBot{
		Key: "",
	}
	param := v20210726.Parameter{
		"msgtype": message.TextStr,
		"text": v20210726.Parameter{
			"content": "测试",
		},
	}
	send, err := bot.Send(param)
	log.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
