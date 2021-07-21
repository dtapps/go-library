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
	if err != nil {
		log.Printf("错误：%v\n", err)
		return
	}
	log.Printf("返回：%s\n", send)
}
