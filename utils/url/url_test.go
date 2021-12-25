package url

import (
	"log"
	"testing"
)

func TestLenCode(t *testing.T) {
	log.Println(LenCode("https://www.dtapp.net"))
}

func TestDeCode(t *testing.T) {
	log.Println(DeCode(LenCode("https://www.dtapp.net")))
}
