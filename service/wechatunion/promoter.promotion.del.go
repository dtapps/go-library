package wechatunion

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newPromotionDelResult(result PromotionDelResponse, body []byte, http gorequest.Response) *PromotionDelResult {
	return &PromotionDelResult{Result: result, Body: body, Http: http}
}

// PromotionDel 删除某个推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_3-%E7%BC%96%E8%BE%91%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (c *Client) PromotionDel(ctx context.Context, promotionSourcePid, promotionSourceName string, notMustParams ...gorequest.Params) (*PromotionDelResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("promotionSourcePid", promotionSourcePid)   // 推广位PID
	params.Set("promotionSourceName", promotionSourceName) // 推广位名称
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/promoter/promotion/del?access_token%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newPromotionDelResult(PromotionDelResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PromotionDelResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPromotionDelResult(response, request.ResponseBody, request), err
}
