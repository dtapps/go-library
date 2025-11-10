package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetJumpDomainConfirmFileResponse struct {
	APIResponse        // 错误
	FileName    string `json:"file_name"`    // 文件名
	FileContent string `json:"file_content"` // 文件内容
}

// GetJumpDomainConfirmFile 获取业务域名校验文件
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/getJumpDomainConfirmFile.html
func (c *Client) GetJumpDomainConfirmFile(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response GetJumpDomainConfirmFileResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_webviewdomain_confirmfile?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetJumpDomainConfirmFileErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 44002:
		return "POST 的数据包为空。post请求body参数不能为空。"
	default:
		return errmsg
	}
}
