package uuid_test

import (
	"fmt"
	"gitee.com/dtapps/go-library/helper/uuid"
	"testing"
)

func TestName(t *testing.T) {
	genUUID := uuid.GenUUID()
	fmt.Println("Hello World")
	fmt.Println(genUUID)
}
