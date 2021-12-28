package dingtalk

import (
	"fmt"
	"net/http"
	"time"
)

type RobotSendResult struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// RobotSend 自定义机器人
// https://open.dingtalk.com/document/group/custom-robot-access
func (app *App) RobotSend(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6
	// 请求
	body, err = app.request(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", app.AccessToken, timestamp, app.sign(timestamp)), params, http.MethodPost)
	return body, err
}
