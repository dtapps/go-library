package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetTesterResponse struct {
	APIResponse // 错误
	Members     []struct {
		Userstr string `json:"userstr"` // 人员对应的唯一字符串
	} `json:"members"` // 人员信息列表
}

// GetTester 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/member-management/getTester.html
func (c *Client) GetTester(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response GetTesterResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "get_experiencer")

	// 请求
	err = c.request(ctx, "wxa/memberauth?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
