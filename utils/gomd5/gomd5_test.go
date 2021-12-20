package gomd5

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(GetMD5Encode("测试"))
	fmt.Println(ToUpper(GetMD5Encode("测试")))
	fmt.Println(ToUpper("测试"))
}
