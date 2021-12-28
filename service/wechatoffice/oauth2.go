package wechatoffice

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println(fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?%s#wechat_redirect", param.Encode()))
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?%s#wechat_redirect", param.Encode())
}

// Oauth2AccessTokenResult 返回参数
type Oauth2AccessTokenResult struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

// Oauth2AccessToken 通过code换取网页授权access_token
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) Oauth2AccessToken(code string) (result Oauth2AccessTokenResult, err error) {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", app.AppId, app.AppSecret, code), map[string]interface{}{}, "GET")
	if err != nil {
		return result, err
	}
	// 解析
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

// Oauth2UserinfoResult 返回参数
type Oauth2UserinfoResult struct {
	Openid     string   `json:"openid"`            // 用户的唯一标识
	Nickname   string   `json:"nickname"`          // 用户昵称
	Sex        int      `json:"sex"`               // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string   `json:"province"`          // 用户个人资料填写的省份
	City       string   `json:"city"`              // 普通用户个人资料填写的城市
	Country    string   `json:"country"`           // 国家，如中国为CN
	Headimgurl string   `json:"headimgurl"`        // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege  []string `json:"privilege"`         // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Unionid    string   `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

// Oauth2Userinfo 拉取用户信息(需scope为 snsapi_userinfo)
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) Oauth2Userinfo(accessToken, openid string) (result Oauth2UserinfoResult, err error) {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid), map[string]interface{}{}, "GET")
	if err != nil {
		return result, err
	}
	// 解析
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
