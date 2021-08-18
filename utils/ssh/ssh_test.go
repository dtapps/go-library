package ssh

import (
	"testing"
)

func TestName(t *testing.T) {

}

func client() {
	Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
