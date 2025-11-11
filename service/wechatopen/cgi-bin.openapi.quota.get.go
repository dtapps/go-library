package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type cgiBinOpenapiQuotaGetResponse struct {
	APIResponse // 错误
	Quota       struct {
		DailyLimit int `json:"daily_limit"` // 当天该账号可调用该接口的次数
		Used       int `json:"used"`        // 当天已经调用的次数
		Remain     int `json:"remain"`      // 当天剩余调用次数
	} `json:"quota"` // quota详情
	RateLimit struct {
		CallCount     int64 `json:"call_count"`     // 周期内可调用数量，单位 次
		RefreshSecond int64 `json:"refresh_second"` // 更新周期，单位 秒`
	} `json:"rate_limit"` // 普通调用频率限制
	ComponentRateLimit struct {
		CallCount     int64 `json:"call_count"`     // 周期内可调用数量，单位 次
		RefreshSecond int64 `json:"refresh_second"` // 更新周期，单位 秒
	} `json:"component_rate_limit"` // 代调用频率限制
}

// CgiBinOpenapiQuotaGet 查询API调用额度
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/getApiQuota.html
func (c *Client) CgiBinOpenapiQuotaGet(ctx context.Context, cgiPath string, notMustParams ...*gorequest.Params) (response cgiBinOpenapiQuotaGetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cgi_path", cgiPath)

	// 请求
	err = c.request(ctx, "cgi-bin/openapi/quota/get?access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetCgiBinOpenapiQuotaGetErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 76021:
		return "cgi_path填错了"
	case 76022:
		return "当前调用接口使用的token与api所属账号不符，详情可看注意事项的说明"
	default:
		return errmsg
	}
}
