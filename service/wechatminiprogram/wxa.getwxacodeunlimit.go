package wechatminiprogram

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostorage"
	"net/http"
)

type WxaGetWxaCodeUnLimitResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type WxaGetWxaCodeUnLimitResult struct {
	Result WxaGetWxaCodeUnLimitResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func newWxaGetWxaCodeUnLimitResult(result WxaGetWxaCodeUnLimitResponse, body []byte, http gorequest.Response, err error) *WxaGetWxaCodeUnLimitResult {
	return &WxaGetWxaCodeUnLimitResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetWxaCodeUnLimit 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) WxaGetWxaCodeUnLimit(ctx context.Context, notMustParams ...gorequest.Params) *WxaGetWxaCodeUnLimitResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/getwxacodeunlimit?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response WxaGetWxaCodeUnLimitResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaGetWxaCodeUnLimitResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetWxaCodeUnLimitResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45009:
		return "调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成"
	case 41030:
		return "page 不合法（页面不存在或者小程序没有发布、根路径前加 /或者携带参数）"
	case 40097:
		return "env_version 不合法"
	}
	return "系统繁忙"
}

// Check 检查
func (resp *WxaGetWxaCodeUnLimitResult) Check() error {
	// 返回是二进制图片，或者json错误
	if resp.Http.ResponseHeader.Get("Content-Type") == "image/jpeg" || resp.Http.ResponseHeader.Get("Content-Type") == "image/png" {
		return nil
	}
	return errors.New("返回不是二进制图片")
}

// Update 上传
func (resp *WxaGetWxaCodeUnLimitResult) Update(storage *gostorage.AliYun, filePath, fileName string) (gostorage.FileInfo, error) {
	return storage.PutObject(bytes.NewReader(resp.Body), filePath, fileName)
}
