package wechatopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetAuthorizerListResponse struct {
	TotalCount int `json:"total_count"` // 授权的账号总数
	List       []struct {
		AuthorizerAppid string `json:"authorizer_appid"` // 已授权账号的 appid
		RefreshToken    string `json:"refresh_token"`    // 刷新令牌authorizer_refresh_token
		AuthTime        int    `json:"auth_time"`        // 授权的时间
	} `json:"list"` // 当前查询的账号基本信息列表
}

type GetAuthorizerListResult struct {
	Result GetAuthorizerListResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newGetAuthorizerListResult(result GetAuthorizerListResponse, body []byte, http gorequest.Response) *GetAuthorizerListResult {
	return &GetAuthorizerListResult{Result: result, Body: body, Http: http}
}

// GetAuthorizerList 拉取已授权的账号信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/authorization-management/getAuthorizerList.html
func (c *Client) GetAuthorizerList(ctx context.Context, authorizerAppid, componentAccessToken string, notMustParams ...gorequest.Params) (*GetAuthorizerListResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/component/api_get_authorizer_list")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId()) // 第三方平台appid

	// 请求
	var response GetAuthorizerListResponse
	request, err := c.request(ctx, span, fmt.Sprintf("cgi-bin/component/api_get_authorizer_list?access_token=%s", componentAccessToken), params, http.MethodPost, &response)
	return newGetAuthorizerListResult(response, request.ResponseBody, request), err
}
