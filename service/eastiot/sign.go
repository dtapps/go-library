package eastiot

import (
	"encoding/json"
	"fmt"
	"gopkg.in/dtapps/go-library.v2/utils/gomd5"
	"sort"
	"strconv"
)

func (app *App) getSign(ApiKey string, p map[string]interface{}) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, app.getString(p[key]))
	}
	signStr += fmt.Sprintf("apiKey=%s", ApiKey)
	return gomd5.ToUpper(signStr)
}

func (app *App) getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
