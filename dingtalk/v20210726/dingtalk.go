package v20210726

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	params "github.com/dtapps/go-library/params/v20210726"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const api = "https://oapi.dingtalk.com/robot/send"

// Parameter 参数
type Parameter map[string]interface{}

// ParameterEncode 参数
type ParameterEncode []string

type DingBot struct {
	Secret      string
	AccessToken string
}

type response struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (bot *DingBot) Send(param Parameter) (response, error) {
	timestamp := time.Now().UnixNano() / 1e6
	var response response
	signStr := sign(timestamp, bot.Secret)
	dingUrl := fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s", api, bot.AccessToken, timestamp, signStr)
	resp, e := http.Post(dingUrl, "application/json", strings.NewReader(param.getRequestData()))
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}
	return response, nil
}

func sign(t int64, secret string) string {
	secStr := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}

// 获取请求数据
func (p Parameter) getRequestData() string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range p {
		args.Set(key, params.GetParamsString(val))
	}
	return args.Encode()
}
