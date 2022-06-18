package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                                  // 错误
}

func NewCgiBinWxOpenQrCodeJumpDownloadResult(result CgiBinWxOpenQrCodeJumpDownloadResponse, body []byte, http gorequest.Response, err error) *CgiBinWxOpenQrCodeJumpDownloadResult {
	return &CgiBinWxOpenQrCodeJumpDownloadResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxOpenQrCodeJumpDownload 获取校验文件名称及内容
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpdownload.html
func (c *Client) CgiBinWxOpenQrCodeJumpDownload() *CgiBinWxOpenQrCodeJumpDownloadResult {
	// 参数
	params := NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdownload?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinWxOpenQrCodeJumpDownloadResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinWxOpenQrCodeJumpDownloadResult(response, request.ResponseBody, request, err)
}
