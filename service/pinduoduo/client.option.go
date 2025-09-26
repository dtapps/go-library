package pinduoduo

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	clientId         string   // POP分配给应用的client_id
	clientSecret     string   // POP分配给应用的client_secret
	mediaId          string   // 媒体ID
	pid              string   // 推广位
	accessToken      string   // 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
	accessTokenScope []string // 授权范围
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

// WithResty 设置 RestyClient
func WithRestyClient(restyClient *resty.Client) Option {
	return Option{F: func(o *Options) {
		o.restyClient = restyClient
	}}
}

// WithResty 设置 restyLog
func WithRestyLog(restyLog *resty_log.Logger) Option {
	return Option{F: func(o *Options) {
		o.restyLog = restyLog
	}}
}

// WithResty 设置 debug
func WithDebug(debug bool) Option {
	return Option{F: func(o *Options) {
		o.debug = debug
	}}
}

// WithResty 设置 clientId
func WithClientId(clientId string) Option {
	return Option{F: func(o *Options) {
		o.clientId = clientId
	}}
}

// WithResty 设置 clientSecret
func WithClientSecret(clientSecret string) Option {
	return Option{F: func(o *Options) {
		o.clientSecret = clientSecret
	}}
}

// WithResty 设置 mediaId
func WithMediaId(mediaId string) Option {
	return Option{F: func(o *Options) {
		o.mediaId = mediaId
	}}
}

// WithResty 设置 pid
func WithPid(pid string) Option {
	return Option{F: func(o *Options) {
		o.pid = pid
	}}
}

// WithResty 设置 accessToken
func WithAccessToken(accessToken string) Option {
	return Option{F: func(o *Options) {
		o.accessToken = accessToken
	}}
}

// WithResty 设置 accessTokenScope
func WithAccessTokenScope(accessTokenScope []string) Option {
	return Option{F: func(o *Options) {
		o.accessTokenScope = accessTokenScope
	}}
}
