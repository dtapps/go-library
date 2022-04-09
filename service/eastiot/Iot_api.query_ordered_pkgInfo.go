package eastiot

import (
	"encoding/json"
	"net/http"
)

type IotApiQueryOrderedPkgInfoResponse struct {
	Code   int `json:"code"`
	Istest int `json:"istest"`
	Data   []struct {
		Name      string  `json:"name"`      // 流量包名字
		PkgId     int64   `json:"pkgId"`     // 流量包ID
		Traffic   int     `json:"traffic"`   // 流量大小，单位:MB
		Ntraffic  float64 `json:"ntraffic"`  // 已用量，单位:MB
		Starttime int     `json:"starttime"` // 流量生效起始时间时间戳
		Endtime   int     `json:"endtime"`   // 流量生效结束时间时间戳
		Addtime   int64   `json:"addtime"`   // 订购时间时间戳
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryOrderedPkgInfoResult struct {
	Result IotApiQueryOrderedPkgInfoResponse // 结果
	Body   []byte                            // 内容
	Err    error                             // 错误
}

func NewIotApiQueryOrderedPkgInfoResult(result IotApiQueryOrderedPkgInfoResponse, body []byte, err error) *IotApiQueryOrderedPkgInfoResult {
	return &IotApiQueryOrderedPkgInfoResult{Result: result, Body: body, Err: err}
}

// IotApiQueryOrderedPkgInfo 查询流量卡已订购流量包
// https://www.showdoc.com.cn/916774523755909/5092045889939625
func (app *App) IotApiQueryOrderedPkgInfo(simId string) *IotApiQueryOrderedPkgInfoResult {
	// 参数
	param := NewParams()
	param.Set("simId", simId)
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("http://m2m.eastiot.net/Api/IotApi/queryOrderedPkgInfo", params, http.MethodPost)
	// 定义
	var response IotApiQueryOrderedPkgInfoResponse
	err = json.Unmarshal(body, &response)
	return NewIotApiQueryOrderedPkgInfoResult(response, body, err)
}
