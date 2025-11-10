package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// Commit 上传代码并生成体验版
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/commit.html
func (c *Client) Commit(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/commit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetCommitErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
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
	default:
		return errmsg
	}
}
