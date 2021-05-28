package uuid_test

import (
	"GoLibrary/helper/uuid"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	genUUID := uuid.GenUUID()
	fmt.Println("Hello World")
	fmt.Println(genUUID)
}
