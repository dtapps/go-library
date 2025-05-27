package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type TimeGet struct {
	TimeGetResponse struct {
		Time      string `json:"time"`
		RequestId string `json:"request_id"`
	} `json:"time_get_response"`
}

// TimeGet 获取拼多多系统时间
// https://open.pinduoduo.com/application/document/api?id=pdd.time.get
func (c *Client) TimeGet(ctx context.Context, notMustParams ...*gorequest.Params) (response TimeGet, err error) {

	// 参数
	params := NewParamsWithType("pdd.time.get", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
