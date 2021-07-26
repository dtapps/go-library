package dmd5

import (
	"fmt"
	v20210726 "github.com/dtapps/go-library/dmd5/v20210726"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(v20210726.Md5("测试"))
}
