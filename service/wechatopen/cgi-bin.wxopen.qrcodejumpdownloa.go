package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
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
func (c *Client) CgiBinWxOpenQrCodeJumpDownload(ctx context.Context) (*CgiBinWxOpenQrCodeJumpDownloadResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/wxopen/qrcodejumpdownload?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinWxOpenQrCodeJumpDownloadResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinWxOpenQrCodeJumpDownloadResult(response, request.ResponseBody, request), nil
}
