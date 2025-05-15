package framework

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 获取字符串类型环境变量，带默认值
func GetEnvDefault(key, defVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defVal
}

// 获取整型环境变量，带默认值
func GetEnvDefaultInt(key string, defVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defVal
}

// 获取布尔型环境变量，带默认值
func GetEnvDefaultBool(key string, defVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defVal
}

// 打印所有环境变量（仅调试用）
func PrintAllEnvs() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

// 获取指定前缀的所有环境变量
func GetEnvsByPrefix(prefix string) map[string]string {
	result := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 && strings.HasPrefix(parts[0], prefix) {
			result[parts[0]] = parts[1]
		}
	}
	return result
}
