package wechatminiprogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostorage"
	"net/http"
)

type WxaGetWxaCodeResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type WxaGetWxaCodeResult struct {
	Result WxaGetWxaCodeResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newWxaGetWxaCodeResult(result WxaGetWxaCodeResponse, body []byte, http gorequest.Response, err error) *WxaGetWxaCodeResult {
	return &WxaGetWxaCodeResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetWxaCode 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (c *Client) WxaGetWxaCode(notMustParams ...gorequest.Params) *WxaGetWxaCodeResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/getwxacode?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaGetWxaCodeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaGetWxaCodeResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetWxaCodeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45029:
		return "生成码个数总和到达最大个数限制"
	case 40097:
		return "env_version 不合法"
	}
	return "系统繁忙"
}

// Check 检查
func (resp *WxaGetWxaCodeResult) Check() error {
	// 返回是二进制图片，或者json错误
	if resp.Http.ResponseHeader.Get("Content-Type") == "image/jpeg" || resp.Http.ResponseHeader.Get("Content-Type") == "image/png" {
		return nil
	}
	return errors.New("返回不是二进制图片")
}

// Update 上传
func (resp *WxaGetWxaCodeResult) Update(storage *gostorage.AliYun, filePath, fileName string) (gostorage.FileInfo, error) {
	return storage.PutObject(bytes.NewReader(resp.Body), filePath, fileName)
}
