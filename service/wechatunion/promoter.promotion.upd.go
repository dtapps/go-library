package wechatunion

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newPromotionUpdResult(result PromotionUpdResponse, body []byte, http gorequest.Response) *PromotionUpdResult {
	return &PromotionUpdResult{Result: result, Body: body, Http: http}
}

// PromotionUpd 编辑推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_3-%E7%BC%96%E8%BE%91%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionUpd(ctx context.Context, notMustParams ...gorequest.Params) (*PromotionUpdResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/promoter/promotion/upd?access_token%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newPromotionUpdResult(PromotionUpdResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PromotionUpdResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPromotionUpdResult(response, request.ResponseBody, request), err
}
