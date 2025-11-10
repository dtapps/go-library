package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetPenaltyListResponse struct {
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

// GetPenaltyList 获取小程序交易体验分违规记录
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/transaction-guarantee/GetPenaltyList.html
func (c *Client) GetPenaltyList(ctx context.Context, authorizerAccessToken string, offset int64, limit int64, notMustParams ...*gorequest.Params) (response GetPenaltyListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("offset", offset) // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过总记录数
	params.Set("limit", limit)   // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过 100

	// 请求
	err = c.request(ctx, "wxaapi/wxamptrade/get_penalty_list?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
