package dingtalk

import (
	"fmt"
	msgtype2 "github.com/dtapps/go-library/dingtalk/msgtype"
	"testing"
)

func TestName(t *testing.T) {
	bot := DingBot{
		Secret:      "",
		AccessToken: "",
	}
	param := Parameter{
		"msgtype": msgtype2.TextStr,
		"text": Parameter{
			"content": "测试",
		},
	}
	send, err := bot.Send(param)
	fmt.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
