package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetIllegalRecordsResponse struct {
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

type GetIllegalRecordsResult struct {
	Result GetIllegalRecordsResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newGetIllegalRecordsResult(result GetIllegalRecordsResponse, body []byte, http gorequest.Response) *GetIllegalRecordsResult {
	return &GetIllegalRecordsResult{Result: result, Body: body, Http: http}
}

// GetIllegalRecords 获取小程序违规处罚记录
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record-management/getIllegalRecords.html
func (c *Client) GetIllegalRecords(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*GetIllegalRecordsResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetIllegalRecordsResponse
	request, err := c.request(ctx, "wxa/getillegalrecords?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetIllegalRecordsResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetIllegalRecordsResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 43007:
		return "检查授权关系"
	default:
		return resp.Result.Errmsg
	}
}
