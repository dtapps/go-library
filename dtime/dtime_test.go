package dtime

import (
	"fmt"
	v20210726 "github.com/dtapps/go-library/dtime/v20210726"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(v20210726.GetCurrentDate())
	fmt.Println(v20210726.GetCurrentUnix())
	fmt.Println(v20210726.GetCurrentMilliUnix())
	fmt.Println(v20210726.GetCurrentNanoUnix())
	fmt.Println(v20210726.GetCurrentWjDate())
}
