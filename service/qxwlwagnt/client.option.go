package qxwlwagnt

import (
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	debug       bool

	baseURL   string // 接口地址
	userName  string // userName
	appKey    string // appKey
	appSecret string // appSecret
}

type Option struct {
	F func(o *Options)
}

func NewOptions(opts []Option) *Options {
	options := &Options{
		restyClient: nil,
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

// WithUserName 设置 userName
func WithUserName(userName string) Option {
	return Option{F: func(o *Options) {
		o.userName = userName
	}}
}

// WithAppKey 设置 appKey
func WithAppKey(appKey string) Option {
	return Option{F: func(o *Options) {
		o.appKey = appKey
	}}
}

// WithAppSecret 设置 appSecret
func WithAppSecret(appSecret string) Option {
	return Option{F: func(o *Options) {
		o.appSecret = appSecret
	}}
}
