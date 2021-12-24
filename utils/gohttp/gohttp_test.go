package gohttp

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	get, err := Get("https://api.dtapp.net/", nil)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", get)
}
