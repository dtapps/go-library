package dingtalk

import (
	v20210726 "github.com/dtapps/go-library/dingtalk/v20210726"
	"github.com/dtapps/go-library/dingtalk/v20210726/msgtype"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bot := v20210726.DingBot{
		Secret:      "",
		AccessToken: "",
	}
	param := v20210726.Parameter{
		"msgtype": msgtype.TextStr,
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
