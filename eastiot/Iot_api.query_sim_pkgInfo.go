package eastiot

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type IotApiQuerySimPkgInfoResponse struct {
	Code   int `json:"code"`
	Istest int `json:"istest"`
	Data   []struct {
		PkgId   int     `json:"pkgId"`   // 流量包ID
		PkgName string  `json:"pkgName"` // 流量包名字
		Price   float64 `json:"price"`   // 流量包成本价格，单位: 元
		Sprice  float64 `json:"sprice"`  // 流量包零售价格，单位: 元
		Traffic int     `json:"traffic"` // 流量包大小，单位: MB
		Type    int     `json:"type"`    // 流量包类型，1:叠加包 2:单月套餐 3:季度套餐 4:半年套餐 5:全年套餐 6:每月套餐 (3个月) 7:每月套餐(6个月) 8:每月套餐(12个月) 0:N天套餐
		Isdm    int     `json:"isdm"`    // 是否依赖主套餐，此字段只有套餐类型为叠加包时有效； 1:依赖主套餐 0:独立
		Isnm    int     `json:"isnm"`    // 是否支持次月生效，此字段只有套餐类型为独立叠加包时有效； 1:支持 0:不支持
		Istest  int     `json:"istest"`  // 是否为体验包； 1:是 0:否
		Isimm   int     `json:"isimm"`   // 订购后是否立即叠加生效； 1:是 0:否
		Stime   string  `json:"stime"`   // 套餐的生效起始日期
		Etime   string  `json:"etime"`   // 套餐的生效结束日期
		Daynum  int     `json:"daynum"`  // 当type=0时，表示套餐有效天数；当type=8 且 daynum>0 时，表示套餐的有效年数
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQuerySimPkgInfoResult struct {
	Result IotApiQuerySimPkgInfoResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func NewIotApiQuerySimPkgInfoResult(result IotApiQuerySimPkgInfoResponse, body []byte, http gorequest.Response, err error) *IotApiQuerySimPkgInfoResult {
	return &IotApiQuerySimPkgInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiQuerySimPkgInfo 流量卡可用流量包查询
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (app *App) IotApiQuerySimPkgInfo(simId string, sd int) *IotApiQuerySimPkgInfoResult {
	// 参数
	param := NewParams()
	param.Set("simId", simId)
	param.Set("sd", sd)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("http://m2m.eastiot.net/Api/IotApi/querySimPkgInfo", params, http.MethodPost)
	// 定义
	var response IotApiQuerySimPkgInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewIotApiQuerySimPkgInfoResult(response, request.ResponseBody, request, err)
}
