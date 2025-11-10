package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaGetUserRiskRankResponse struct {
	APIResponse       // 错误
	RiskRank    int   `json:"risk_rank"` // 用户风险等级，合法值为0,1,2,3,4，数字越大风险越高。
	UnoinId     int64 `json:"unoin_id"`  // 唯一请求标识，标记单次请求
}

// WxaGetUserRiskRank 获取用户安全等级
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/safety-control-capability/getUserRiskRank.html
func (c *Client) WxaGetUserRiskRank(ctx context.Context, notMustParams ...*gorequest.Params) (response WxaGetUserRiskRankResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAuthorizerAppid())

	// 请求
	err = c.request(ctx, "wxa/getuserriskrank?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
