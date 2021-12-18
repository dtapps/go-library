package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/service/dingtalk/config"
	"github.com/dtapps/go-library/service/dingtalk/message"
	"github.com/dtapps/go-library/utils/gojson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type DingBot struct {
	Secret      string
	AccessToken string
}

type Response struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (bot *DingBot) Send(msg message.Message) (Response, error) {
	timestamp := time.Now().UnixNano() / 1e6
	var response Response
	signStr := sign(timestamp, bot.Secret)
	dingUrl := fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s", config.Api, bot.AccessToken, timestamp, signStr)
	toString, err := gojson.MarshalToString(msg)
	if err != nil {
		return response, err
	}
	resp, e := http.Post(dingUrl, "application/json", strings.NewReader(toString))
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
