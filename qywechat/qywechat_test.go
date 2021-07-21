package qywechat

import (
	"gopkg.in/dtapps/go-library.v2/qywechat/message"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bot := QyBot{
		Key: "",
	}
	msg := message.Message{
		MsgType: message.TextStr,
		Text: message.Text_{
			Content: "测试",
		},
	}
	send, err := bot.Send(msg)
	log.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
