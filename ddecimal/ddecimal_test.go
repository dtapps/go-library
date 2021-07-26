package ddecimal

import (
	v20210726 "github.com/dtapps/go-library/daes/v20210726"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.Println(v20210726.Decimal(2.3333))
}
