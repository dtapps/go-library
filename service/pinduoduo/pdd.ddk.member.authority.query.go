package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type MemberAuthorityQuery struct {
	AuthorityQueryResponse struct {
		Bind      int64  `json:"bind"`
		RequestId string `json:"request_id"`
	} `json:"authority_query_response"`
}

// MemberAuthorityQuery 查询是否绑定备案
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.search
func (c *Client) MemberAuthorityQuery(ctx context.Context, notMustParams ...*gorequest.Params) (response MemberAuthorityQuery, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.member.authority.query", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	err = c.request(ctx, params, &response)
	return
}
