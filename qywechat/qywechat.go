package qywechat

import (
	"encoding/json"
	"fmt"
	params2 "github.com/dtapps/go-library/params"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const api = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"

// Parameter 参数
type Parameter map[string]interface{}

// ParameterEncode 参数
type ParameterEncode []string

type QyBot struct {
	Key string
}

type response struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (bot *QyBot) Send(param Parameter) (response, error) {
	var response response
	qyUrl := fmt.Sprintf("%s?key=%s", api, bot.Key)
	resp, e := http.Post(qyUrl, "application/json", strings.NewReader(param.getRequestData()))
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

// 获取请求数据
func (p Parameter) getRequestData() string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range p {
		args.Set(key, params2.GetParamsString(val))
	}
	return args.Encode()
}
