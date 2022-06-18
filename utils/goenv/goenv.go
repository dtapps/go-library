package goenv

import (
	"go.dtapp.net/library/utils/gostring"
	"os"
	"strings"
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

func GetEnvs(key string) string {
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			println(parts[0], parts[1])
			if parts[0] == key {
				return parts[1]
			}
		}
	}
	return ""
}
