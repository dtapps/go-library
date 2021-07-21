package dingtalk

import (
	"gopkg.in/dtapps/go-library.v2/dingtalk/message"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bot := DingBot{
		Secret:      "",
		AccessToken: "",
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
