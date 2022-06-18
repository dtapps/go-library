package dingdanxia

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanPrivilegeResponse struct {
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

type WaiMaiMeituanPrivilegeResult struct {
	Result WaiMaiMeituanPrivilegeResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
	Err    error                          // 错误
}

func NewWaiMaiMeituanPrivilegeResult(result WaiMaiMeituanPrivilegeResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanPrivilegeResult {
	return &WaiMaiMeituanPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanPrivilege 美团外卖CPS推广API接口
// https://www.dingdanxia.com/doc/174/173
func (app *App) WaiMaiMeituanPrivilege(sid string, generateWeApp, qrcode bool) *WaiMaiMeituanPrivilegeResult {
	// 参数
	param := NewParams()
	param.Set("sid", sid)
	param.Set("generate_we_app", generateWeApp)
	param.Set("qrcode", qrcode)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://api.tbk.dingdanxia.com/waimai/meituan_privilege", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanPrivilegeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWaiMaiMeituanPrivilegeResult(response, request.ResponseBody, request, err)
}
