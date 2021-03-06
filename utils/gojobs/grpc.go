package gojobs

import (
	"fmt"
	"github.com/dtapps/go-library/utils/goip"
)

var ip string

func configIp() {
	ip = goip.GetOutsideIp()
}

const prefix = "cron:"

const prefixIp = "cron_%s:"

func prefixSprintf(str string) string {
	return fmt.Sprintf(prefixIp, str)
}
