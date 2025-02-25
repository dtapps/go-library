package meituan

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type MediaGetReferralLinkResponse struct {
	Status    int    `json:"status"`              // 响应码，0成功，其他失败
	Message   string `json:"message"`             // 响应文案
	Data      string `json:"data,omitempty"`      // 返回对应的推广链接，这里的链接才能实现跟单计佣
	SkuViewId string `json:"skuViewId,omitempty"` // 若用text进行入参取链，且返回的推广链接为商品券链接，则返回对应商品的展示ID，可以根据该ID查商品券接口获取对应的展示信息和佣金信息
}

type MediaGetReferralLinkResult struct {
	Result MediaGetReferralLinkResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newMediaGetReferralLinkResult(result MediaGetReferralLinkResponse, body []byte, http gorequest.Response) *MediaGetReferralLinkResult {
	return &MediaGetReferralLinkResult{Result: result, Body: body, Http: http}
}

// MediaGetReferralLink 获取推广链接接口
// 支持获取活动物料、到店/到家/买菜业务类型的推广链接；支持按活动物料ID、商品券展示ID、目标链接的形式获取对应的推广链接；支持appkey-sid两级渠道追踪推广效果。需要用POST方法调用接口。
// https://media.meituan.com/pc/index.html#/materials/api-detail/get_referral_link
func (c *MediaClient) MediaGetReferralLink(ctx context.Context, notMustParams ...*gorequest.Params) (*MediaGetReferralLinkResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response MediaGetReferralLinkResponse
	request, err := c.request(ctx, "cps_open/common/api/v1/get_referral_link", http.MethodPost, params, &response)
	return newMediaGetReferralLinkResult(response, request.ResponseBody, request), err
}
