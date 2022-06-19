package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PromotionDelResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type PromotionDelResult struct {
	Result PromotionDelResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func newPromotionDelResult(result PromotionDelResponse, body []byte, http gorequest.Response, err error) *PromotionDelResult {
	return &PromotionDelResult{Result: result, Body: body, Http: http, Err: err}
}

// PromotionDel 删除某个推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_3-%E7%BC%96%E8%BE%91%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionDel(promotionSourcePid, promotionSourceName string) *PromotionDelResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("promotionSourcePid", promotionSourcePid)   // 推广位PID
	params.Set("promotionSourceName", promotionSourceName) // 推广位名称
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/promoter/promotion/del?access_token%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response PromotionDelResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPromotionDelResult(response, request.ResponseBody, request, err)
}
