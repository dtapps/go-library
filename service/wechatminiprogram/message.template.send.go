package wechatminiprogram

import (
	"fmt"
)

func (app *App) MessageTemplateSend(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", app.AccessToken), params, "POST")
	return
}
