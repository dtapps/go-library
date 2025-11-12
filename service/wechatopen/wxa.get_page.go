package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetCodePageResponse struct {
	APIResponse          // 错误
	PageList    []string `json:"page_list"` // page_list 页面配置列表
}

// GetCodePage 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/code-management/api_getcodepage.html
func (c *Client) GetCodePage(ctx context.Context, notMustParams ...*gorequest.Params) (response GetCodePageResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_page?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}
