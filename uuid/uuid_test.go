package uuid_test

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v2/uuid"
	"testing"
)

func TestName(t *testing.T) {
	genUUID := uuid.GenUUID()
	fmt.Println("Hello World")
	fmt.Println(genUUID)
}
