package pinduoduo

import (
	v20210726 "github.com/dtapps/go-library/pinduoduo/v20210726"
	_type "github.com/dtapps/go-library/pinduoduo/v20210726/type"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	duo := v20210726.PinDuoDuo{
		ClientId:     "",
		ClientSecret: "",
	}
	param := v20210726.Parameter{
		"keyword": "小米",
	}
	send, err := duo.Send(_type.GoodsSearch, param)
	log.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("错误：%v", err)
		return
	}
}
