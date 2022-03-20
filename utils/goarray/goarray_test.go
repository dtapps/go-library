package goarray

import "testing"

func TestName(t *testing.T) {
	ss := []string{"c2:37:7d:2b:78:c3", "c2:37:7d:2b:78:c4", "14:98:77:5a:58:66", "c2:37:7d:2b:78:a3", "c2:37:7d:2b:78:a4", "36:ce:4f:d2:84:00", "36:ce:4f:d2:84:04", "36:98:77:50:03:1b", "14:98:77:50:03:1b", "36:ce:4f:d2:84:00", "e2:ed:05:15:9c:2a", "e2:ed:05:15:9c:2a"}
	t.Log(TurnString(ss))
}
