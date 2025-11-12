package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetIllegalRecordsResponse struct {
	APIResponse // 错误
	Records     []struct {
		IllegalRecordId string `json:"illegal_record_id"` // 违规处罚记录id
		CreateTime      int    `json:"create_time"`       // 违规处罚时间
		IllegalReason   string `json:"illegal_reason"`    // 违规原因
		IllegalContent  string `json:"illegal_content"`   // 违规内容
		RuleUrl         string `json:"rule_url"`          // 规则链接
		RuleName        string `json:"rule_name"`         // 违反的规则名称
	} `json:"records"` // 违规处罚记录列表
}

// GetIllegalRecords 获取小程序违规处罚记录
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/record-management/api_getillegalrecords.html
func (c *Client) GetIllegalRecords(ctx context.Context, notMustParams ...*gorequest.Params) (response GetIllegalRecordsResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("wxa/getillegalrecords"), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetIllegalRecordsErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 43007:
		return "检查授权关系"
	default:
		return errmsg
	}
}
