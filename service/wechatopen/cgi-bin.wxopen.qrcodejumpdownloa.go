package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpDownloadResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type CgiBinWxOpenQrCodeJumpDownloadResult struct {
	Result CgiBinWxOpenQrCodeJumpDownloadResponse // 结果
	Body   []byte                                 // 内容
	Http   gorequest.Response                     // 请求
}

func newCgiBinWxOpenQrCodeJumpDownloadResult(result CgiBinWxOpenQrCodeJumpDownloadResponse, body []byte, http gorequest.Response) *CgiBinWxOpenQrCodeJumpDownloadResult {
	return &CgiBinWxOpenQrCodeJumpDownloadResult{Result: result, Body: body, Http: http}
}

// CgiBinWxOpenQrCodeJumpDownload 获取校验文件名称及内容
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpdownload.html
func (c *Client) CgiBinWxOpenQrCodeJumpDownload(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*CgiBinWxOpenQrCodeJumpDownloadResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/wxopen/qrcodejumpdownload?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newCgiBinWxOpenQrCodeJumpDownloadResult(CgiBinWxOpenQrCodeJumpDownloadResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinWxOpenQrCodeJumpDownloadResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinWxOpenQrCodeJumpDownloadResult(response, request.ResponseBody, request), err
}
