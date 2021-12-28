package ip

import (
	"net"
	"testing"
)

var app App

func TestIp(t *testing.T) {
	t.Log(net.ParseIP("61.241.55.180").To4())
	t.Log(net.ParseIP("240e:3b4:38e4:3295:7093:af6c:e545:f2e9").To16())
	t.Log(app.V4db.Find("61.241.55.180"))
	t.Log(app.V6db.Find("240e:3b4:38e4:3295:7093:af6c:e545:f2e9"))
	t.Log(app.V4Region.MemorySearch("61.241.55.180"))
}
