package dingtalk

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type RobotSendResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type RobotSendResult struct {
	Result RobotSendResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRobotSendResult(result RobotSendResponse, body []byte, http gorequest.Response) *RobotSendResult {
	return &RobotSendResult{Result: result, Body: body, Http: http}
}

// RobotSend 自定义机器人
// https://open.dingtalk.com/document/group/custom-robot-access
func (c *Client) RobotSend(ctx context.Context, notMustParams ...*gorequest.Params) (*RobotSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/robot/send?access_token=%s&timestamp=%d&sign=%s", c.GetAccessToken(), timestamp, c.sign(timestamp)), params, http.MethodPost)
	if err != nil {
		return newRobotSendResult(RobotSendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RobotSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRobotSendResult(response, request.ResponseBody, request), err
}
