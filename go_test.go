package go_library

import (
	"gopkg.in/dtapps/go-library.v2/library"
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	log.Println(library.Version())
}
