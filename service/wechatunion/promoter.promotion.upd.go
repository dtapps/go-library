package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PromotionUpdResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type PromotionUpdResult struct {
	Result PromotionUpdResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func newPromotionUpdResult(result PromotionUpdResponse, body []byte, http gorequest.Response, err error) *PromotionUpdResult {
	return &PromotionUpdResult{Result: result, Body: body, Http: http, Err: err}
}

// PromotionUpd 编辑推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_3-%E7%BC%96%E8%BE%91%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionUpd(notMustParams ...gorequest.Params) *PromotionUpdResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/promoter/promotion/upd?access_token%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response PromotionUpdResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPromotionUpdResult(response, request.ResponseBody, request, err)
}
