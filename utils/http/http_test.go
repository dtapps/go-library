package http

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(GetJson("https://api.dtapp.net/", "", nil))
}
