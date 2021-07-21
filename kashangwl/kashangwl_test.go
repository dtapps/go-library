package kashangwl

import (
	"gopkg.in/dtapps/go-library.v2/kashangwl/message"
	"gopkg.in/dtapps/go-library.v2/kashangwl/url"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	wl := KaShangWl{
		CustomerId:  0000000,
		CustomerKey: "xxx",
	}
	msg := message.Order{
		OrderId: 827669582783,
	}
	send, err := wl.Send(msg, url.Order)
	log.Printf("send：%s\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
