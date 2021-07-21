package pinduoduo

import (
	_type "gopkg.in/dtapps/go-library.v2/pinduoduo/type"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	duo := PinDuoDuo{
		ClientId:     "",
		ClientSecret: "",
	}
	param := Parameter{
		"keyword": "小米",
	}
	send, err := duo.Send(_type.GoodsSearch, param)
	log.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("错误：%v", err)
		return
	}
}
