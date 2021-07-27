package qywechat

import (
	"fmt"
	message2 "github.com/dtapps/go-library/service/qywechat/message"
	"testing"
)

func TestName(t *testing.T) {
	bot := QyBot{
		Key: "",
	}
	msg := message2.Message{
		MsgType: message2.TextStr,
		Text: message2.Text_{
			Content: "测试",
		},
	}
	send, err := bot.Send(msg)
	fmt.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
