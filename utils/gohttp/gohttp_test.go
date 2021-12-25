package gohttp

import (
	"testing"
)

func TestName(t *testing.T) {
	get, err := Get("https://api.dtapp.net/", nil)
	t.Logf("%+v\n", get)
	if err != nil {
		return
	}
}
