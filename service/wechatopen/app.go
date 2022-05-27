package wechatopen

import (
	"encoding/json"
	"errors"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gotime"
	"gorm.io/gorm"
	"net/http"
)

// App 微信公众号服务
type App struct {
	componentAccessToken  string // 第三方平台 access_token
	componentVerifyTicket string // 微信后台推送的 ticket
	preAuthCode           string // 预授权码

	authorizerAccessToken  string // 接口调用令牌
	authorizerRefreshToken string // 刷新令牌
	AuthorizerAppid        string // 授权方 appid

	ComponentAppId     string // 第三方平台 appid
	ComponentAppSecret string // 第三方平台 app_secret
	MessageToken       string
	MessageKey         string

	Mongo gomongo.App // 非关系数据库服务
	Db    *gorm.DB    // 关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		get, err := gohttp.Get(url, params)
		// 日志
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}

// GetAuthorizerAccessToken 获取授权方令牌
func (app *App) GetAuthorizerAccessToken() string {
	if app.Db == nil {
		return app.authorizerAccessToken
	}
	var result AuthorizerAccessToken
	app.Db.Where("component_app_id = ?", app.ComponentAppId).Where("authorizer_app_id = ?", app.AuthorizerAppid).Where("expire_time >= ?", gotime.Current().Format()).Last(&result)
	return result.AuthorizerAccessToken
}

// GetAuthorizerRefreshToken 获取刷新令牌
func (app *App) GetAuthorizerRefreshToken() string {
	if app.Db == nil {
		return app.authorizerRefreshToken
	}
	var result AuthorizerAccessToken
	app.Db.Where("component_app_id = ?", app.ComponentAppId).Where("authorizer_app_id = ?", app.AuthorizerAppid).Last(&result)
	return result.AuthorizerRefreshToken
}

// GetPreAuthCode 获取预授权码
func (app *App) GetPreAuthCode() string {
	if app.Db == nil {
		return app.preAuthCode
	}
	var result PreAuthCode
	app.Db.Where("app_id = ?", app.ComponentAppId).Where("expire_time >= ?", gotime.Current().Format()).Last(&result)
	return result.PreAuthCode
}

// GetComponentAccessToken 获取 access_token
func (app *App) GetComponentAccessToken() string {
	if app.Db == nil {
		return app.componentAccessToken
	}
	var result ComponentAccessToken
	app.Db.Where("app_id = ?", app.ComponentAppId).Where("expire_time >= ?", gotime.Current().Format()).Last(&result)
	return result.ComponentAccessToken
}

// GetComponentVerifyTicket 获取 Ticket
func (app *App) GetComponentVerifyTicket() string {
	if app.Db == nil {
		return app.componentVerifyTicket
	}
	var result ComponentVerifyTicket
	app.Db.Where("app_id = ?", app.ComponentAppId).Where("expire_time >= ?", gotime.Current().Format()).Last(&result)
	return result.ComponentVerifyTicket
}
