package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type CgiBinOpenSameEnTityResponse struct {
	APIResponse      // 错误
	SameEntity  bool `json:"same_entity"` // 是否同主体；true表示同主体；false表示不同主体
}

// CgiBinOpenSameEnTity 获取授权绑定的商户号列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/cloudbase-common/wechatpay/getWechatPayList.html
func (c *Client) CgiBinOpenSameEnTity(ctx context.Context, notMustParams ...*gorequest.Params) (response CgiBinOpenSameEnTityResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/open/sameentity?access_token="+c.GetComponentAccessToken(), params, http.MethodGet, &response)
	return
}
