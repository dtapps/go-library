package dhttp

import (
	"fmt"
	v20210726 "github.com/dtapps/go-library/dhttp/v20210726"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(v20210726.GetJson("https://api.dtapp.net/", "", nil))
}
