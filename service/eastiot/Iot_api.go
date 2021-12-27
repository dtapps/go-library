package eastiot

import "net/http"

// IotApiQueryUserBalanceResult 返回参数
type IotApiQueryUserBalanceResult struct {
	Code int `json:"code"`
	Data struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// IotApiQueryUserBalance 余额查询
// https://www.showdoc.com.cn/916774523755909/4857910459512420
func (app *App) IotApiQueryUserBalance() (body []byte, err error) {
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/queryUserBalance", map[string]interface{}{}, http.MethodPost)
	return body, err
}

type IotApiGetAllSimTypeResult struct {
	Code int `json:"code"`
	Data []struct {
		Type   int    `json:"type"`   // 卡类型
		Name   string `json:"name"`   // 类型名
		MOrder int    `json:"mOrder"` // 是否支持单次充值多个流量包，0:不支持 1:支持
	} `json:"data"`
	Msg string `json:"msg"`
}

// IotApiGetAllSimType 卡类型列表查询
// https://www.showdoc.com.cn/916774523755909/4858492092033167
func (app *App) IotApiGetAllSimType() (body []byte, err error) {
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/getAllSimType", map[string]interface{}{}, http.MethodPost)
	return body, err
}

type IotApiQueryUserPkgInfoResult struct {
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

// IotApiQueryUserPkgInfo 账户可用流量包查询
// https://www.showdoc.com.cn/916774523755909/4850094776758927
func (app *App) IotApiQueryUserPkgInfo() (body []byte, err error) {
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/queryUserPkgInfo", map[string]interface{}{}, http.MethodPost)
	return body, err
}

type IotApiRechargeSimResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// IotApiRechargeSim 单卡流量充值
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (app *App) IotApiRechargeSim(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/rechargeSim", params, http.MethodPost)
	return body, err
}

type IotApiQuerySimPkgInfoResult struct {
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

// IotApiQuerySimPkgInfo 流量卡可用流量包查询
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (app *App) IotApiQuerySimPkgInfo(simId string, sd int) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("simId", simId)
	param.Set("sd", sd)
	params := app.NewParamsWith(param)
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/querySimPkgInfo", params, http.MethodPost)
	return body, err
}

type IotApiQueryOrderedPkgInfoResult struct {
	Code   int `json:"code"`
	Istest int `json:"istest"`
	Data   []struct {
		Name      string  `json:"name"`      // 流量包名字
		PkgId     int64   `json:"pkgId"`     // 流量包ID
		Traffic   int     `json:"traffic"`   // 流量大小，单位:MB
		Ntraffic  float64 `json:"ntraffic"`  // 已用量，单位:MB
		Starttime int     `json:"starttime"` // 流量生效起始时间时间戳
		Endtime   int     `json:"endtime"`   // 流量生效结束时间时间戳
		Addtime   int     `json:"addtime"`   // 订购时间时间戳
	} `json:"data"`
	Msg string `json:"msg"`
}

// IotApiQueryOrderedPkgInfo 查询流量卡已订购流量包
// https://www.showdoc.com.cn/916774523755909/5092045889939625
func (app *App) IotApiQueryOrderedPkgInfo(simId string) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("simId", simId)
	params := app.NewParamsWith(param)
	// 请求
	body, err = app.request("http://m2m.eastiot.net/Api/IotApi/queryOrderedPkgInfo", params, http.MethodPost)
	return body, err
}
