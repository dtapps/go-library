package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkTimeGetResponse struct {
	TimeGetResponse struct {
		Time      string `json:"time"`
		RequestId string `json:"request_id"`
	} `json:"time_get_response"`
}

type PddDdkTimeGetResult struct {
	Result PddDdkTimeGetResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newPddDdkTimeGetResult(result PddDdkTimeGetResponse, body []byte, http gorequest.Response) *PddDdkTimeGetResult {
	return &PddDdkTimeGetResult{Result: result, Body: body, Http: http}
}

// TimeGet 获取拼多多系统时间
// https://open.pinduoduo.com/application/document/api?id=pdd.time.get
func (c *Client) TimeGet(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkTimeGetResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.time.get")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.time.get", notMustParams...)

	// 请求
	var response PddDdkTimeGetResponse
	request, err := c.request(ctx, span, params, &response)
	return newPddDdkTimeGetResult(response, request.ResponseBody, request), err
}
