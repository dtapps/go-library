package durl

import (
	v20210726 "github.com/dtapps/go-library/durl/v20210726"
	"log"
	"testing"
)

func TestLenCode(t *testing.T) {
	u := "https://www.dtapp.net"
	log.Println(v20210726.LenCode(u))
}

func TestDeCode(t *testing.T) {
	u := "https%3A%2F%2Fwww.dtapp.net"
	log.Println(v20210726.DeCode(u))
}

func TestParseQuery(t *testing.T) {
	u := "https://api.dtapp.net/v8/ip/qqwry?ip=113.118.170.199"
	log.Println(v20210726.ParseQuery(u))
}
