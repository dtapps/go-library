package dingdanxia

import (
	"dtapps/dta/library/utils/gohttp"
	"encoding/json"
	"net/http"
)

type WaimaiMeituanPrivilegeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		H5        string `json:"h5"`
		ShortH5   string `json:"short_h5"`
		Deeplink  string `json:"deeplink"`
		H5Evoke   string `json:"h5_evoke"`
		Tkl       string `json:"tkl"`
		WeAppInfo struct {
			AppId    string `json:"app_id"`
			PagePath string `json:"page_path"`
			MiniCode string `json:"miniCode"`
		} `json:"we_app_info"`
		Qrcode string `json:"qrcode"`
	} `json:"data"`
}

type WaimaiMeituanPrivilegeResult struct {
	Result WaimaiMeituanPrivilegeResponse // 结果
	Body   []byte                         // 内容
	Http   gohttp.Response                // 请求
	Err    error                          // 错误
}

func NewWaimaiMeituanPrivilegeResult(result WaimaiMeituanPrivilegeResponse, body []byte, http gohttp.Response, err error) *WaimaiMeituanPrivilegeResult {
	return &WaimaiMeituanPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaimaiMeituanPrivilege 美团外卖CPS推广API接口
// https://www.dingdanxia.com/doc/174/173
func (app *App) WaimaiMeituanPrivilege(sid string, generateWeApp, qrcode bool) *WaimaiMeituanPrivilegeResult {
	// 参数
	param := NewParams()
	param.Set("sid", sid)
	param.Set("generate_we_app", generateWeApp)
	param.Set("qrcode", qrcode)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://api.tbk.dingdanxia.com/waimai/meituan_privilege", params, http.MethodPost)
	// 定义
	var response WaimaiMeituanPrivilegeResponse
	err = json.Unmarshal(request.Body, &response)
	return NewWaimaiMeituanPrivilegeResult(response, request.Body, request, err)
}
