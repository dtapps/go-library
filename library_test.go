package go_library

import (
	"fmt"
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println(Version())
	log.Println(Version())
}
