package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type MemberAuthorityQueryResponse struct {
	AuthorityQueryResponse struct {
		Bind      int64  `json:"bind"`
		RequestId string `json:"request_id"`
	} `json:"authority_query_response"`
}

type MemberAuthorityQueryResult struct {
	Result MemberAuthorityQueryResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newMemberAuthorityQueryResult(result MemberAuthorityQueryResponse, body []byte, http gorequest.Response) *MemberAuthorityQueryResult {
	return &MemberAuthorityQueryResult{Result: result, Body: body, Http: http}
}

// MemberAuthorityQuery 查询是否绑定备案
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.search
func (c *Client) MemberAuthorityQuery(ctx context.Context, notMustParams ...gorequest.Params) (*MemberAuthorityQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.member.authority.query")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.member.authority.query", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	var response MemberAuthorityQueryResponse
	request, err := c.request(ctx, span, params, &response)
	return newMemberAuthorityQueryResult(response, request.ResponseBody, request), err
}
