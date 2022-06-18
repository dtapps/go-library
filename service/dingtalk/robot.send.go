package dingtalk

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error              // 错误
}

func NewRobotSendResult(result RobotSendResponse, body []byte, http gorequest.Response, err error) *RobotSendResult {
	return &RobotSendResult{Result: result, Body: body, Http: http, Err: err}
}

// RobotSend 自定义机器人
// https://open.dingtalk.com/document/group/custom-robot-access
func (app *App) RobotSend(notMustParams ...Params) *RobotSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6
	// 请求
	request, err := app.request(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", app.accessToken, timestamp, app.sign(timestamp)), params, http.MethodPost)
	// 定义
	var response RobotSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRobotSendResult(response, request.ResponseBody, request, err)
}
