package go_library

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v2/library"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(library.Version())
}
