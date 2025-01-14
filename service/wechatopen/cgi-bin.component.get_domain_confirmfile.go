package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetThirdpartyJumpDomainConfirmFileResponse struct {
	Errcode     int    `json:"errcode"`      // 返回码
	Errmsg      string `json:"errmsg"`       // 返回码信息
	FileName    string `json:"file_name"`    // 文件名
	FileContent string `json:"file_content"` // 文件内容
}

type GetThirdpartyJumpDomainConfirmFileResult struct {
	Result GetThirdpartyJumpDomainConfirmFileResponse // 结果
	Body   []byte                                     // 内容
	Http   gorequest.Response                         // 请求
}

func newGetThirdpartyJumpDomainConfirmFileResult(result GetThirdpartyJumpDomainConfirmFileResponse, body []byte, http gorequest.Response) *GetThirdpartyJumpDomainConfirmFileResult {
	return &GetThirdpartyJumpDomainConfirmFileResult{Result: result, Body: body, Http: http}
}

// GetThirdpartyJumpDomainConfirmFile 获取第三方平台业务域名校验文件
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/getThirdpartyJumpDomainConfirmFile.html
func (c *Client) GetThirdpartyJumpDomainConfirmFile(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*GetThirdpartyJumpDomainConfirmFileResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetThirdpartyJumpDomainConfirmFileResponse
	request, err := c.request(ctx, "cgi-bin/component/get_domain_confirmfile?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetThirdpartyJumpDomainConfirmFileResult(response, request.ResponseBody, request), err
}
