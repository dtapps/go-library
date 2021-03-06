package dingdanxia

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanSgPrivilegeResponse struct {
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

type WaiMaiMeituanSgPrivilegeResult struct {
	Result WaiMaiMeituanSgPrivilegeResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
	Err    error                            // 错误
}

func newWaiMaiMeituanSgPrivilegeResult(result WaiMaiMeituanSgPrivilegeResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanSgPrivilegeResult {
	return &WaiMaiMeituanSgPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanSgPrivilege 美团闪购CPS推广API接口
// https://www.dingdanxia.com/doc/195/173
func (c *Client) WaiMaiMeituanSgPrivilege(sid string, generateWeApp, qrcode bool) *WaiMaiMeituanSgPrivilegeResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("sid", sid)
	param.Set("generate_we_app", generateWeApp)
	param.Set("qrcode", qrcode)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/waimai/meituan_sg_privilege", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanSgPrivilegeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanSgPrivilegeResult(response, request.ResponseBody, request, err)
}
