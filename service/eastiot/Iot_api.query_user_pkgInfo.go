package eastiot

import (
	"encoding/json"
	"net/http"
)

type IotApiQueryUserPkgInfoResponse struct {
	Code int `json:"code"`
	Data []struct {
		Type    int     `json:"type"`
		PkgId   int64   `json:"pkgId"`
		PkgName string  `json:"pkgName"`
		Price   float64 `json:"price"`
		Sprice  float64 `json:"sprice"`
		Traffic int     `json:"traffic"`
		Caltype int     `json:"caltype"`
		SimType int     `json:"simType"`
		Isdm    int     `json:"isdm"`
		Isnm    int     `json:"isnm"`
		Istest  int     `json:"istest"`
		Isimm   int     `json:"isimm"`
		Daynum  int     `json:"daynum"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryUserPkgInfoResult struct {
	Result IotApiQueryUserPkgInfoResponse // 结果
	Body   []byte                         // 内容
	Err    error                          // 错误
}

func NewIotApiQueryUserPkgInfoResult(result IotApiQueryUserPkgInfoResponse, body []byte, err error) *IotApiQueryUserPkgInfoResult {
	return &IotApiQueryUserPkgInfoResult{Result: result, Body: body, Err: err}
}

// IotApiQueryUserPkgInfo 账户可用流量包查询
// https://www.showdoc.com.cn/916774523755909/4850094776758927
func (app *App) IotApiQueryUserPkgInfo() *IotApiQueryUserPkgInfoResult {
	// 请求
	body, err := app.request("http://m2m.eastiot.net/Api/IotApi/queryUserPkgInfo", map[string]interface{}{}, http.MethodPost)
	// 定义
	var response IotApiQueryUserPkgInfoResponse
	err = json.Unmarshal(body, &response)
	return NewIotApiQueryUserPkgInfoResult(response, body, err)
}
