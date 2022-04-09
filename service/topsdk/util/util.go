package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

type LocalTime time.Time

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func UnmarshalJSON(data []byte, v interface{}) (err error) {
	return json.Unmarshal(data, v)
}

func ConvertStructList(data interface{}) string {
	if data == nil {
		return "[]"
	}
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func ConvertStruct(data interface{}) string {
	if data == nil {
		return "{}"
	}
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func ConvertBasicList(data interface{}) string {
	if data == nil {
		return "[]"
	}
	return strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", ",", -1)
}

func HandleJsonResponse(jsonStr string, v interface{}) (err error) {

	if strings.Contains(jsonStr[0:20], "error_response") {
		err := &TopApiRequestError{}
		jsonStr = jsonStr[18 : len(jsonStr)-1]
		err2 := json.Unmarshal([]byte(jsonStr), err)
		if err2 != nil {
			return err2
		}
		return err
	}
	return json.Unmarshal([]byte(jsonStr), v)
}

func GetSign(publicParam map[string]interface{}, data map[string]interface{}, secret string) string {
	var allParamMap = make(map[string]interface{})
	for k, v := range data {
		allParamMap[k] = v
	}
	for k, v := range publicParam {
		allParamMap[k] = v
	}
	var keyList []string
	for k := range allParamMap {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	var signStr = ""
	for _, key := range keyList {
		value := allParamMap[key]
		signStr = signStr + fmt.Sprintf("%v%v", key, value)
		//if(value != ""){
		//	signStr = signStr + fmt.Sprintf("%v%v", key, value)
		//}
	}
	fmt.Println(signStr)
	sign := strings.ToUpper(hmacSha256(signStr, secret))
	return sign
}

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
