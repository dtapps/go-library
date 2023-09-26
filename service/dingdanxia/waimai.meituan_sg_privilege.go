package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanSgPrivilegeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		H5        string `json:"h5"`       // H5 领券
		Deeplink  string `json:"deeplink"` // Deeplink领券
		H5Evoke   string `json:"h5_evoke"` // H5 内唤起页
		ShortH5   string `json:"short_h5"` // h5短连接
		WeAppInfo struct {
			AppId    string `json:"app_id"`    // 小程序ID
			PagePath string `json:"page_path"` // 小程序路径
			MiniCode string `json:"miniCode"`  // 小程序码
		} `json:"we_app_info"` // 小程序信息
		Qrcode string `json:"qrcode"` // 海报
	} `json:"data"`
}

type WaiMaiMeituanSgPrivilegeResult struct {
	Result WaiMaiMeituanSgPrivilegeResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newWaiMaiMeituanSgPrivilegeResult(result WaiMaiMeituanSgPrivilegeResponse, body []byte, http gorequest.Response) *WaiMaiMeituanSgPrivilegeResult {
	return &WaiMaiMeituanSgPrivilegeResult{Result: result, Body: body, Http: http}
}

// WaiMaiMeituanSgPrivilege 美团闪购CPS推广API接口
// https://www.dingdanxia.com/doc/195/173
func (c *Client) WaiMaiMeituanSgPrivilege(ctx context.Context, sid string, generateWeApp, qrcode bool, notMustParams ...gorequest.Params) (*WaiMaiMeituanSgPrivilegeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sid", sid)                       // 渠道方用户唯一标识,渠道可自定义,长度不超过50，参数中不能包含dingdanxia，用于向用户返佣,支持小写字母和数字的格式,其它字符可能造成无法正常跟单
	params.Set("generate_we_app", generateWeApp) // 是否生成小程序推广信息
	params.Set("qrcode", qrcode)                 // 是否生成二维码海报
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_sg_privilege", params, http.MethodPost)
	if err != nil {
		return newWaiMaiMeituanSgPrivilegeResult(WaiMaiMeituanSgPrivilegeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WaiMaiMeituanSgPrivilegeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanSgPrivilegeResult(response, request.ResponseBody, request), err
}
