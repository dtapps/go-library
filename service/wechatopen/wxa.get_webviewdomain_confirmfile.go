package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetJumpDomainConfirmFileResponse struct {
	Errcode     int    `json:"errcode"`      // 返回码
	Errmsg      string `json:"errmsg"`       // 返回码信息
	FileName    string `json:"file_name"`    // 文件名
	FileContent string `json:"file_content"` // 文件内容
}

type GetJumpDomainConfirmFileResult struct {
	Result GetJumpDomainConfirmFileResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newGetJumpDomainConfirmFileResult(result GetJumpDomainConfirmFileResponse, body []byte, http gorequest.Response) *GetJumpDomainConfirmFileResult {
	return &GetJumpDomainConfirmFileResult{Result: result, Body: body, Http: http}
}

// GetJumpDomainConfirmFile 获取业务域名校验文件
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/getJumpDomainConfirmFile.html
func (c *Client) GetJumpDomainConfirmFile(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*GetJumpDomainConfirmFileResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetJumpDomainConfirmFileResponse
	request, err := c.request(ctx, "wxa/get_webviewdomain_confirmfile?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetJumpDomainConfirmFileResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetJumpDomainConfirmFileResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 44002:
		return "POST 的数据包为空。post请求body参数不能为空。"
	default:
		return resp.Result.Errmsg
	}
}
