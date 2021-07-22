package kashangwl

import (
	_url "gopkg.in/dtapps/go-library.v2/kashangwl/url"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	wl := KaShangWl{
		CustomerId:  0,
		CustomerKey: "",
	}

	param := Parameter{
		"order_id": 827669582783,
	}
	send, err := wl.Send(_url.Order, param)
	log.Printf("send：%s\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
