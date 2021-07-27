package djson

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
