package goenv

import (
	"os"
	"strconv"
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
		i, _ := strconv.Atoi(val)
		return i
	}
	return defVal
}

func GetEnvDefaultBool(key string, defVal bool) bool {
	val, ok := os.LookupEnv(key)
	if ok {
		b, _ := strconv.ParseBool(val)
		return b
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
