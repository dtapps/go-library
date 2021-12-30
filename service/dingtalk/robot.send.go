package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RobotSendResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type RobotSendResult struct {
	Result RobotSendResponse // 结果
	Body   []byte            // 内容
	Err    error             // 错误
}

func NewRobotSendResult(result RobotSendResponse, body []byte, err error) *RobotSendResult {
	return &RobotSendResult{Result: result, Body: body, Err: err}
}

// RobotSend 自定义机器人
// https://open.dingtalk.com/document/group/custom-robot-access
func (app *App) RobotSend(notMustParams ...Params) *RobotSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6
	// 请求
	body, err := app.request(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", app.AccessToken, timestamp, app.sign(timestamp)), params, http.MethodPost)
	// 定义
	var response RobotSendResponse
	err = json.Unmarshal(body, &response)
	return NewRobotSendResult(response, body, err)
}
