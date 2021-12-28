package wechatoffice

import (
	"fmt"
	"net/http"
)

// MessageTemplateSend 模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (app *App) MessageTemplateSend(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", app.AccessToken), params, http.MethodPost)
	return
}
