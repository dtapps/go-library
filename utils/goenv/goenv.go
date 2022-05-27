package goenv

import (
	"go.dtapp.net/library/utils/gostring"
	"os"
)

func GetEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defVal
}

func GetEnvDefaultInt(key string, defVal int) int {
	val, ok := os.LookupEnv(key)
	if ok {
		return gostring.ToInt(val)
	}
	return defVal
}
