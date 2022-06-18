package gitee

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"net/http"
	//	"strconv"
)

// App 基本配置
type App struct {
	ClientID     string
	ClientSecret string
	RedirectUri  string
	AccessToken  string
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	// 请求参数
	marshal, _ := json.Marshal(params)
	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	var response *http.Response
	response, err = httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	// 	处理成功
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)

	return resp, err
}
