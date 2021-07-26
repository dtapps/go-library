package durl

import (
	"log"
	"testing"
)

func TestLenCode(t *testing.T) {
	u := "https://www.dtapp.net"
	log.Println(LenCode(u))
}

func TestDeCode(t *testing.T) {
	u := "https%3A%2F%2Fwww.dtapp.net"
	log.Println(DeCode(u))
}

func TestParseQuery(t *testing.T) {
	u := "https://api.dtapp.net/v8/ip/qqwry?ip=113.118.170.199"
	log.Println(ParseQuery(u))
}
