package go_library

import (
	_ "go.dtapp.net/library/service"
	_ "go.dtapp.net/library/utils"
)

func Version() string {
	return "v1.0.45"
}
