package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetThirdpartyJumpDomainConfirmFileResponse struct {
	APIResponse        // 错误
	FileName    string `json:"file_name"`    // 文件名
	FileContent string `json:"file_content"` // 文件内容
}

// GetThirdpartyJumpDomainConfirmFile 获取第三方平台业务域名校验文件
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/getThirdpartyJumpDomainConfirmFile.html
func (c *Client) GetThirdpartyJumpDomainConfirmFile(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response GetThirdpartyJumpDomainConfirmFileResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/component/get_domain_confirmfile?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
