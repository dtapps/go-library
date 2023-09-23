package wechatunion

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PromotionAddResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Pid     string `json:"pid"`     // 推广位ID，PID
}

type PromotionAddResult struct {
	Result PromotionAddResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newPromotionAddResult(result PromotionAddResponse, body []byte, http gorequest.Response) *PromotionAddResult {
	return &PromotionAddResult{Result: result, Body: body, Http: http}
}

// PromotionAdd 添加推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_1-%E6%B7%BB%E5%8A%A0%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionAdd(ctx context.Context, promotionSourceName string, notMustParams ...*gorequest.Params) (*PromotionAddResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("promotionSourceName", promotionSourceName) // 推广位名称
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/promoter/promotion/add?access_token%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newPromotionAddResult(PromotionAddResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PromotionAddResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPromotionAddResult(response, request.ResponseBody, request), err
}
