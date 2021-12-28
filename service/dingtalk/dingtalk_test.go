package dingtalk

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v2/service/dingtalk/message"
	"testing"
)

func TestName(t *testing.T) {
	bot := DingBot{
		Secret:      "gSEC05342ba24a7eb2e1dbeae61b3df997eb1a97b7cda414566876e983f1db0fec0b",
		AccessToken: "caad81de7f6b218bb7d085264d4885714c805cc80a460690a0d19db91a05dd174",
	}
	msg := message.Message{
		MsgType: message.TextStr,
		Text: message.Text_{
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
