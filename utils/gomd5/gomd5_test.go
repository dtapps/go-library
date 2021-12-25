package gomd5

import (
	"testing"
)

func TestMd5(t *testing.T) {
	t.Logf(GetMD5Encode("测试"))
	t.Logf(ToUpper(GetMD5Encode("测试")))
	t.Logf(ToUpper("测试"))
}
