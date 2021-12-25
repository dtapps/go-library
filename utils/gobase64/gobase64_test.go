package gobase64

import (
	"testing"
)

func TestEncode(t *testing.T) {
	t.Log(Encode("广东茂名"))
}

func TestDecode(t *testing.T) {
	t.Log(Decode(Encode("广东茂名")))
}
