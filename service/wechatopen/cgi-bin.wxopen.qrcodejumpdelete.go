package wechatopen

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpDeleteResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinWxOpenQrCodeJumpDeleteResult struct {
	Result CgiBinWxOpenQrCodeJumpDeleteResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
	Err    error                                // 错误
}

func newCgiBinWxOpenQrCodeJumpDeleteResult(result CgiBinWxOpenQrCodeJumpDeleteResponse, body []byte, http gorequest.Response, err error) *CgiBinWxOpenQrCodeJumpDeleteResult {
	return &CgiBinWxOpenQrCodeJumpDeleteResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxOpenQrCodeJumpDelete 删除已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpdelete.html
func (c *Client) CgiBinWxOpenQrCodeJumpDelete(prefix string) *CgiBinWxOpenQrCodeJumpDeleteResult {
	// 参数
	params := gorequest.NewParams()
	params["prefix"] = prefix
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/cgi-bin/wxopen/qrcodejumpdelete?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinWxOpenQrCodeJumpDeleteResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinWxOpenQrCodeJumpDeleteResult(response, request.ResponseBody, request, err)
}
