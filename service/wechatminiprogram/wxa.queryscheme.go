package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WxaQueryScheme 查询小程序 scheme 码，及长期有效 quota
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html
func (app *App) WxaQueryScheme(notMustParams ...Params) *BusinessGetLiveInfoResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/queryscheme?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response BusinessGetLiveInfoResponse
	err = json.Unmarshal(body, &response)
	return NewBusinessGetLiveInfoResult(response, body, err)
}
