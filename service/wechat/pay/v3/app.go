package v3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type App struct {
	AppId           string // 小程序或者公众号的appid
	AppSecret       string
	MchId           string // 微信支付的商户id
	AesKey          string
	ApiV3           string
	PrivateSerialNo string // 私钥证书号
	MchPrivateKey   string // 路径 apiclient_key.pem
}

// ErrResp 错误返回
type ErrResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  struct {
		Field    string      `json:"field,omitempty"`
		Value    interface{} `json:"value"`
		Issue    string      `json:"issue,omitempty"`
		Location string      `json:"location"`
	} `json:"detail"`
}

func (app *App) request(url string, params map[string]interface{}) (resp []byte, result ErrResp, err error) {

	canonicalURL := fmt.Sprintf("%s/%s", WechatPayAPIServer, url)
	method := "POST"
	authorization, _ := app.authorization(method, params, canonicalURL)

	marshal, _ := json.Marshal(params)

	var req *http.Request
	req, err = http.NewRequest(method, canonicalURL, bytes.NewReader(marshal))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN")
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	httpClient := &http.Client{}
	var response *http.Response
	response, err = httpClient.Do(req)

	if err != nil {
		return nil, result, err
	}

	// 	处理成功
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)

	// 检查错误
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, result, err
	}

	// 检查请求错误
	if response.StatusCode == 200 {
		return resp, result, err
	}

	return nil, result, err
}
