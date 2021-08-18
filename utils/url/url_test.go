package url

import (
	"fmt"
	"testing"
)

func TestLenCode(t *testing.T) {
	u := "https://www.dtapp.net"
	fmt.Println(LenCode(u))
}

func TestDeCode(t *testing.T) {
	u := "https%3A%2F%2Fwww.dtapp.net"
	fmt.Println(DeCode(u))
}

func TestParseQuery(t *testing.T) {
	u := "https://api.dtapp.net/v8/ip/qqwry?ip=113.118.170.199"
	fmt.Println(ParseQuery(u))
}
