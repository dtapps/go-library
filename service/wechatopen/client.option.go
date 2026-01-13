package wechatopen

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.LoggerMiddleware
	debug       bool

	baseURL string // 接口地址

	componentAppId        string // 第三方平台appid
	componentAppSecret    string // 第三方平台app_secret
	messageToken          string // 第三方平台消息令牌
	messageKey            string // 第三方平台消息密钥
	componentAccessToken  string // 第三方平台access_token
	componentVerifyTicket string // 第三方平台推送ticket
	componentPreAuthCode  string // 第三方平台预授权码

	authorizerAppid          string // 授权方appid
	authorizerAccessToken    string // 授权方access_token
	authorizerRefreshToken   string // 授权方refresh_token
	authorizerReleaseVersion string // 授权方release_version
}

type Option struct {
	F func(o *Options)
}

func NewOptions(opts []Option) *Options {
	options := &Options{
		restyClient: nil,
		restyLog:    nil,
	}
	options.Apply(opts)
	return options
}

func (o *Options) Apply(opts []Option) {
	for _, op := range opts {
		op.F(o)
	}
}

// WithRestyClient 设置 RestyClient
func WithRestyClient(restyClient *resty.Client) Option {
	return Option{F: func(o *Options) {
		o.restyClient = restyClient
	}}
}

// WithRestyClientIf 设置 RestyClient
func WithRestyClientIf(enable bool, restyClient *resty.Client) Option {
	return Option{F: func(o *Options) {
		if enable {
			o.restyClient = restyClient
		}
	}}
}

// WithRestyLog 设置 restyLog
func WithRestyLog(restyLog *resty_log.LoggerMiddleware) Option {
	return Option{F: func(o *Options) {
		o.restyLog = restyLog
	}}
}

// WithRestyLogIf 设置 restyLog
func WithRestyLogIf(enable bool, restyLog *resty_log.LoggerMiddleware) Option {
	return Option{F: func(o *Options) {
		if enable {
			o.restyLog = restyLog
		}
	}}
}

// WithResty 设置 debug
func WithDebug(debug bool) Option {
	return Option{F: func(o *Options) {
		o.debug = debug
	}}
}

// WithResty 设置 URL
func WithURL(baseURL string) Option {
	return Option{F: func(o *Options) {
		o.baseURL = baseURL
	}}
}

// WithComponentAppId 设置 componentAppId
func WithComponentAppId(componentAppId string) Option {
	return Option{F: func(o *Options) {
		o.componentAppId = componentAppId
	}}
}

// WithComponentAppSecret 设置 componentAppSecret
func WithComponentAppSecret(componentAppSecret string) Option {
	return Option{F: func(o *Options) {
		o.componentAppSecret = componentAppSecret
	}}
}

// WithMessageToken 设置 messageToken
func WithMessageToken(messageToken string) Option {
	return Option{F: func(o *Options) {
		o.messageToken = messageToken
	}}
}

// WithMessageKey 设置 messageKey
func WithMessageKey(messageKey string) Option {
	return Option{F: func(o *Options) {
		o.messageKey = messageKey
	}}
}

// WithComponentAccessToken 设置 componentAccessToken
func WithComponentAccessToken(componentAccessToken string) Option {
	return Option{F: func(o *Options) {
		o.componentAccessToken = componentAccessToken
	}}
}

// WithComponentVerifyTicket 设置 componentVerifyTicket
func WithComponentVerifyTicket(componentVerifyTicket string) Option {
	return Option{F: func(o *Options) {
		o.componentVerifyTicket = componentVerifyTicket
	}}
}

// WithComponentPreAuthCode 设置 componentPreAuthCode
func WithComponentPreAuthCode(componentPreAuthCode string) Option {
	return Option{F: func(o *Options) {
		o.componentPreAuthCode = componentPreAuthCode
	}}
}

// WithAuthorizerAppid 设置 authorizerAppid
func WithAuthorizerAppid(authorizerAppid string) Option {
	return Option{F: func(o *Options) {
		o.authorizerAppid = authorizerAppid
	}}
}

// WithAuthorizerAccessToken 设置 authorizerAccessToken
func WithAuthorizerAccessToken(authorizerAccessToken string) Option {
	return Option{F: func(o *Options) {
		o.authorizerAccessToken = authorizerAccessToken
	}}
}

// WithAuthorizerRefreshToken 设置 authorizerRefreshToken
func WithAuthorizerRefreshToken(authorizerRefreshToken string) Option {
	return Option{F: func(o *Options) {
		o.authorizerRefreshToken = authorizerRefreshToken
	}}
}

// WithAuthorizerReleaseVersion 设置 authorizerReleaseVersion
func WithAuthorizerReleaseVersion(authorizerReleaseVersion string) Option {
	return Option{F: func(o *Options) {
		o.authorizerReleaseVersion = authorizerReleaseVersion
	}}
}
