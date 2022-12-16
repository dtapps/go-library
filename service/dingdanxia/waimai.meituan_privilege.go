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
	Err    error                          // 错误
}

func newWaiMaiMeituanPrivilegeResult(result WaiMaiMeituanPrivilegeResponse, body []byte, http gorequest.Response, err error) *WaiMaiMeituanPrivilegeResult {
	return &WaiMaiMeituanPrivilegeResult{Result: result, Body: body, Http: http, Err: err}
}

// WaiMaiMeituanPrivilege 美团外卖CPS推广API接口
// https://www.dingdanxia.com/doc/174/173
func (c *Client) WaiMaiMeituanPrivilege(ctx context.Context, sid string, generateWeApp bool, channels int, qrcode bool) *WaiMaiMeituanPrivilegeResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("sid", sid)                       // 渠道方用户唯一标识,渠道可自定义,长度不超过50，参数中不能包含dingdanxia，用于向用户返佣,支持小写字母和数字的格式,其它字符可能造成无法正常跟单
	param.Set("generate_we_app", generateWeApp) // 是否生成小程序推广信息
	param.Set("channels", channels)             // 推广渠道 1-小程序推广,2-公众号推广,3-app推广,4-社群推广 默认1 ，请务必选择对应渠道推广，选择错误会影响佣金比例
	param.Set("qrcode", qrcode)                 // 二维码图片
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_privilege", params, http.MethodPost)
	// 定义
	var response WaiMaiMeituanPrivilegeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanPrivilegeResult(response, request.ResponseBody, request, err)
}
