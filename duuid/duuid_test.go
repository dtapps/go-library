package duuid_test

import (
	v20210726 "github.com/dtapps/go-library/duuid/v20210726"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.Println(v20210726.GenUUID())
}
