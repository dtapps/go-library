package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaCommitResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaCommitResult struct {
	Result WxaCommitResponse // 结果
	Body   []byte            // 内容
	Err    error             // 错误
}

func NewWxaCommitResult(result WxaCommitResponse, body []byte, err error) *WxaCommitResult {
	return &WxaCommitResult{Result: result, Body: body, Err: err}
}

// WxaCommit 上传小程序代码并生成体验版
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html
func (app *App) WxaCommit(notMustParams ...Params) *WxaCommitResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/commit?access_token=%s", app.authorizerAccessToken), params, http.MethodPost)
	// 定义
	var response WxaCommitResponse
	err = json.Unmarshal(body, &response)
	return NewWxaCommitResult(response, body, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaCommitResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85013:
		return "无效的自定义配置"
	case 85014:
		return "无效的模板编号"
	case 85043:
		return "模板错误"
	case 85044:
		return "代码包超过大小限制"
	case 85045:
		return "ext_json 有不存在的路径"
	case 85046:
		return "tabBar 中缺少 path"
	case 85047:
		return "pages 字段为空"
	case 85048:
		return "ext_json 解析失败"
	case 80082:
		return "没有权限使用该插件"
	case 80067:
		return "找不到使用的插件"
	case 80066:
		return "非法的插件版本"
	case 9402202:
		return "请勿频繁提交，待上一次操作完成后再提交"
	case 9402203:
		return `标准模板ext_json错误，传了不合法的参数， 如果是标准模板库的模板，则ext_json支持的参数仅为{"extAppid":'', "ext": {}, "window": {}}`
	}
	return "系统繁忙"
}