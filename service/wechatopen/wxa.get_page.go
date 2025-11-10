package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaGetPageResponse struct {
	APIResponse          // 错误
	PageList    []string `json:"page_list"` // page_list 页面配置列表
}

// WxaGetPage 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (c *Client) WxaGetPage(ctx context.Context, notMustParams ...*gorequest.Params) (response WxaGetPageResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_page?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}
