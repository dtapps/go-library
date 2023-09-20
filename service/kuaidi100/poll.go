package kuaidi100

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PollResponse struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
}

type PollResult struct {
	Result PollResponse       // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newPollResult(result PollResponse, body []byte, http gorequest.Response) *PollResult {
	return &PollResult{Result: result, Body: body, Http: http}
}

// Poll 实时快递查询接口
// https://api.kuaidi100.com/document/5f0ffb5ebc8da837cbd8aefc
func (c *Client) Poll(ctx context.Context, notMustParams ...*gorequest.Params) (*PollResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/poll", params, http.MethodPost)
	if err != nil {
		return newPollResult(PollResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PollResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPollResult(response, request.ResponseBody, request), err
}
