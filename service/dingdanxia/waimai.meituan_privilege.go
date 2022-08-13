package dingdanxia

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
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

func newWaiMaiMeituanPrivilegeResult(result WaiMaiMeituanPrivilegeResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanPrivilegeResult {
	return &WaiMaiMeituanPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanPrivilege 美团外卖CPS推广API接口
// https://www.dingdanxia.com/doc/174/173
func (c *Client) WaiMaiMeituanPrivilege(ctx context.Context, sid string, generateWeApp, qrcode bool) *WaiMaiMeituanPrivilegeResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("sid", sid)
	param.Set("generate_we_app", generateWeApp)
	param.Set("qrcode", qrcode)
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_privilege", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanPrivilegeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanPrivilegeResult(response, request.ResponseBody, request, err)
}
