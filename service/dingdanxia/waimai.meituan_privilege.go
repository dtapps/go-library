package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanPrivilegeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		H5        string `json:"h5"`       // H5 领券
		ShortH5   string `json:"short_h5"` // H5 领券短链接
		Deeplink  string `json:"deeplink"` // Deeplink领券
		H5Evoke   string `json:"h5_evoke"` // H5 内唤起页
		Qrcode    string `json:"qrcode"`   // 二维码海报图片路径
		Tkl       string `json:"tkl"`      // 团口令
		WeAppInfo struct {
			AppId    string `json:"app_id"`    // 小程序ID
			PagePath string `json:"page_path"` // 小程序路径
			MiniCode string `json:"miniCode"`  // 小程序码
		} `json:"we_app_info"` // 小程序信息
	} `json:"data"`
}

type WaiMaiMeituanPrivilegeResult struct {
	Result WaiMaiMeituanPrivilegeResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newWaiMaiMeituanPrivilegeResult(result WaiMaiMeituanPrivilegeResponse, body []byte, http gorequest.Response) *WaiMaiMeituanPrivilegeResult {
	return &WaiMaiMeituanPrivilegeResult{Result: result, Body: body, Http: http}
}

// WaiMaiMeituanPrivilege 美团外卖CPS推广API接口
// https://www.dingdanxia.com/doc/174/173
func (c *Client) WaiMaiMeituanPrivilege(ctx context.Context, notMustParams ...*gorequest.Params) (*WaiMaiMeituanPrivilegeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_privilege", params, http.MethodPost)
	if err != nil {
		return newWaiMaiMeituanPrivilegeResult(WaiMaiMeituanPrivilegeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WaiMaiMeituanPrivilegeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanPrivilegeResult(response, request.ResponseBody, request), err
}
