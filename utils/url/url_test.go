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
