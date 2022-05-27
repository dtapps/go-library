package dingdanxia

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gohttp"
	"net/http"
)

type WaimaiMeituanSgPrivilegeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		H5        string `json:"h5"`
		Deeplink  string `json:"deeplink"`
		H5Evoke   string `json:"h5_evoke"`
		ShortH5   string `json:"short_h5"`
		WeAppInfo struct {
			AppId    string `json:"app_id"`
			PagePath string `json:"page_path"`
			MiniCode string `json:"miniCode"`
		} `json:"we_app_info"`
		Qrcode string `json:"qrcode"`
	} `json:"data"`
}

type WaimaiMeituanSgPrivilegeResult struct {
	Result WaimaiMeituanSgPrivilegeResponse // 结果
	Body   []byte                           // 内容
	Http   gohttp.Response                  // 请求
	Err    error                            // 错误
}

func NewWaimaiMeituanSgPrivilegeResult(result WaimaiMeituanSgPrivilegeResponse, body []byte, http gohttp.Response, err error) *WaimaiMeituanSgPrivilegeResult {
	return &WaimaiMeituanSgPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaimaiMeituanSgPrivilege 美团闪购CPS推广API接口
// https://www.dingdanxia.com/doc/195/173
func (app *App) WaimaiMeituanSgPrivilege(sid string, generateWeApp, qrcode bool) *WaimaiMeituanSgPrivilegeResult {
	// 参数
	param := NewParams()
	param.Set("sid", sid)
	param.Set("generate_we_app", generateWeApp)
	param.Set("qrcode", qrcode)
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request("https://api.tbk.dingdanxia.com/waimai/meituan_sg_privilege", params, http.MethodPost)
	// 定义
	var response WaimaiMeituanSgPrivilegeResponse
	err = json.Unmarshal(request.Body, &response)
	return NewWaimaiMeituanSgPrivilegeResult(response, request.Body, request, err)
}
