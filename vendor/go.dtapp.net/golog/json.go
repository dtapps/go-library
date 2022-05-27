package golog

import "encoding/json"

func JsonEncodeNoError(data interface{}) string {
	jsons, _ := json.Marshal(data)
	return string(jsons)
}
