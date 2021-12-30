package wechatoffice

import (
	"fmt"
	"net/url"
)

// Oauth2 用户同意授权，获取code
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) Oauth2(redirectUri, state string) string {
	param := url.Values{}
	param.Add("appid", app.AppId)          // 公众号的唯一标识
	param.Add("redirect_uri", redirectUri) // 授权后重定向的回调链接地址， 请使用 urlEncode 对链接进行处理
	param.Add("response_type", "code")     // 返回类型
	param.Add("scope", "snsapi_userinfo")  // 应用授权作用域
	param.Add("state", state)              // 重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?%s#wechat_redirect", param.Encode())
}
