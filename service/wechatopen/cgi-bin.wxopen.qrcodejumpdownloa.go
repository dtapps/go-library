package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinWxOpenQrCodeJumpDownloadResponse struct {
	APIResponse        // 错误
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

// CgiBinWxOpenQrCodeJumpDownload 获取校验文件名称及内容
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/downloadQRCodeText.html
func (c *Client) CgiBinWxOpenQrCodeJumpDownload(ctx context.Context, notMustParams ...*gorequest.Params) (response CgiBinWxOpenQrCodeJumpDownloadResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/wxopen/qrcodejumpdownload?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
