package kashangwl

import (
	"fmt"
	_url "github.com/dtapps/go-library/kashangwl/url"
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
	fmt.Printf("send：%s\n", send)
	if err != nil {
		t.Errorf("err：%v\n", err)
		return
	}
}
