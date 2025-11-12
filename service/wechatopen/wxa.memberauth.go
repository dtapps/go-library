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
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/member-management/api_gettester.html
func (c *Client) GetTester(ctx context.Context, notMustParams ...*gorequest.Params) (response GetTesterResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "get_experiencer")

	// 请求
	err = c.request(ctx, "wxa/memberauth?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
