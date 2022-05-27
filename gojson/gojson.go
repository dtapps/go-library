package gojson

import "encoding/json"

func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func MarshalToString(msg interface{}) (string, error) {
	j, e := json.Marshal(msg)
	if e != nil {
		return "", e
	}
	return string(j), nil
}

func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

func JsonDecodeNoError(data string) map[string]interface{} {
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

func JsonEncodeNoError(data interface{}) string {
	jsons, _ := json.Marshal(data)
	return string(jsons)
}

func JsonDecodesNoError(data string) []string {
	var dat []string
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}
