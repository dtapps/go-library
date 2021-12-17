package gouuid_test

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gouuid"
	"testing"
)

func TestGetUuId(t *testing.T) {
	fmt.Println(gouuid.GetUuId())
}
