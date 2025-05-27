package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type GoodsZsUnitUrlGen struct {
	GoodsZsUnitUrlGenResponse struct {
		MobileShortUrl           string `json:"mobile_short_url,omitempty"`             // 对应出参mobile_url的短链接，与mobile_url功能一致。
		MobileUrl                string `json:"mobile_url,omitempty"`                   // 普通长链，微信环境下进入领券页点领券拉起小程序，浏览器环境下直接拉起APP，未安装拼多多APP时落地页点领券拉起登录页
		MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url,omitempty"` // 推广短链接（唤起拼多多app）
		MultiGroupMobileUrl      string `json:"multi_group_mobile_url,omitempty"`       // 推广长链接（可唤起拼多多app）
		MultiGroupShortUrl       string `json:"multi_group_short_url,omitempty"`        // 双人团推广短链接
		MultiGroupUrl            string `json:"multi_group_url,omitempty"`              // 双人团推广长链接
		ShortUrl                 string `json:"short_url,omitempty"`                    // 对应出参url的短链接，与url功能一致。
		Url                      string `json:"url,omitempty"`                          // 普通长链。微信环境下进入领券页点领券拉起小程序，浏览器环境下优先拉起微信小程序
		WeixinShortLink          string `json:"weixin_short_link,omitempty"`            // 小程序短链，点击可直接唤起微信小程序
	} `json:"goods_zs_unit_generate_response"`
}

// GoodsZsUnitUrlGen 多多进宝转链接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.zs.unit.url.gen
func (c *Client) GoodsZsUnitUrlGen(ctx context.Context, notMustParams ...*gorequest.Params) (response GoodsZsUnitUrlGen, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.zs.unit.url.gen", notMustParams...)
	params.Set("p_id", c.GetPid())

	// 请求
	err = c.request(ctx, params, &response)
	return
}
