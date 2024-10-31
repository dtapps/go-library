package pushdeer

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PushResponse struct {
	Code  int `json:"code"`
	Errno int `json:"errno"`
	Data  struct {
		Pushid int `json:"pushid"`
		Meta   struct {
			Android struct {
				MessageIds struct {
					MessageId []string `json:"messageId"`
				} `json:"messageIds"`
				RequestId string   `json:"requestId"`
				Devices   []string `json:"devices"`
			} `json:"android"`
			Ios struct {
				MessageIds struct {
					MessageId []string `json:"messageId"`
				} `json:"messageIds"`
				RequestId string   `json:"requestId"`
				Devices   []string `json:"devices"`
			} `json:"ios"`
			Devices []string `json:"devices"`
		} `json:"meta"`
	} `json:"data,omitempty"`
	Message string `json:"message"`
}

type PushResult struct {
	Result PushResponse       // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newPushResult(result PushResponse, body []byte, http gorequest.Response) *PushResult {
	return &PushResult{Result: result, Body: body, Http: http}
}

// Push 推送消息
// https://doc.sc3.ft07.com/server/api
func (c *Client) Push(ctx context.Context, notMustParams ...gorequest.Params) (*PushResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response PushResponse
	request, err := c.request(ctx, c.config.url, params, http.MethodPost, &response)
	return newPushResult(response, request.ResponseBody, request), err
}
