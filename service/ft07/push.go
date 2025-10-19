package ft07

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
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

// Push 推送消息
// https://doc.sc3.ft07.com/server/api
func (c *Client) Push(ctx context.Context, notMustParams ...*gorequest.Params) (response PushResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, c.config.baseURL, params, http.MethodPost, &response)
	return
}
