package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetPenaltyListResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
	Records []struct {
		IllegalRecordId string `json:"illegal_record_id"` // 违规处罚记录id
		CreateTime      int    `json:"create_time"`       // 违规处罚时间
		IllegalReason   string `json:"illegal_reason"`    // 违规原因
		IllegalContent  string `json:"illegal_content"`   // 违规内容
		RuleUrl         string `json:"rule_url"`          // 规则链接
		RuleName        string `json:"rule_name"`         // 违反的规则名称
	} `json:"records"` // 违规处罚记录列表
}

type GetPenaltyListResult struct {
	Result GetPenaltyListResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newGetPenaltyListResult(result GetPenaltyListResponse, body []byte, http gorequest.Response) *GetPenaltyListResult {
	return &GetPenaltyListResult{Result: result, Body: body, Http: http}
}

// GetPenaltyList 获取小程序交易体验分违规记录
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/transaction-guarantee/GetPenaltyList.html
func (c *Client) GetPenaltyList(ctx context.Context, authorizerAccessToken string, offset int64, limit int64, notMustParams ...*gorequest.Params) (*GetPenaltyListResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("offset", offset) // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过总记录数
	params.Set("limit", limit)   // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过 100

	// 请求
	var response GetPenaltyListResponse
	request, err := c.request(ctx, "wxaapi/wxamptrade/get_penalty_list?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetPenaltyListResult(response, request.ResponseBody, request), err
}
