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
func (c *Client) WaiMaiMeituanSgPrivilege(ctx context.Context, notMustParams ...*gorequest.Params) (*WaiMaiMeituanSgPrivilegeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
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
