package dssh

import (
	v20210726 "gopkg.in/dtapps/go-library.v2/dssh/v20210726"
	"testing"
)

func TestName(t *testing.T) {

}

func client() {
	v20210726.Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
