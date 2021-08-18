package time

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(GetCurrentDate())
	fmt.Println(GetCurrentUnix())
	fmt.Println(GetCurrentMilliUnix())
	fmt.Println(GetCurrentNanoUnix())
	fmt.Println(GetCurrentWjDate())
}
