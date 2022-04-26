package go_library

import (
	_ "github.com/dtapps/go-library/service"
	_ "github.com/dtapps/go-library/utils"
)

func Version() string {
	return "v1.0.45"
}
