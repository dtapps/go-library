package wechatqy

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	appId       string
	agentId     int
	secret      string
	redirectUri string
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

// WithResty 设置 appId
func WithAppId(appId string) Option {
	return Option{F: func(o *Options) {
		o.appId = appId
	}}
}

// WithResty 设置 agentId
func WithAgentId(agentId int) Option {
	return Option{F: func(o *Options) {
		o.agentId = agentId
	}}
}

// WithResty 设置 secret
func WithSecret(secret string) Option {
	return Option{F: func(o *Options) {
		o.secret = secret
	}}
}

// WithResty 设置 redirectUri
func WithRedirectUri(redirectUri string) Option {
	return Option{F: func(o *Options) {
		o.redirectUri = redirectUri
	}}
}
