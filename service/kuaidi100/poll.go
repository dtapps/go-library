package kuaidi100

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PollResponse struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
}

// Poll 实时快递查询接口
// https://api.kuaidi100.com/document/5f0ffb5ebc8da837cbd8aefc
func (c *Client) Poll(ctx context.Context, notMustParams ...*gorequest.Params) (response PollResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "poll", params, http.MethodPost, &response)
	return
}
