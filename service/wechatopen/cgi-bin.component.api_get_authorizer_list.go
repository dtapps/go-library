package wechatopen

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetAuthorizerListResponse struct {
	TotalCount int `json:"total_count"` // 授权的账号总数
	List       []struct {
		AuthorizerAppid string `json:"authorizer_appid"` // 已授权账号的 appid
		RefreshToken    string `json:"refresh_token"`    // 刷新令牌authorizer_refresh_token
		AuthTime        int    `json:"auth_time"`        // 授权的时间
	} `json:"list"` // 当前查询的账号基本信息列表
}

// GetAuthorizerList 拉取已授权的账号信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/authorization-management/getAuthorizerList.html
func (c *Client) GetAuthorizerList(ctx context.Context, notMustParams ...*gorequest.Params) (response GetAuthorizerListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId()) // 第三方平台appid

	// 请求
	err = c.request(ctx, fmt.Sprintf("cgi-bin/component/api_get_authorizer_list?access_token=%s", c.GetComponentAccessToken()), params, http.MethodPost, &response)
	return
}
