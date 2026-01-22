package pushdeer

import (
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	debug       bool

	baseURL string // 接口地址
	pushKey string // 推送key
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

// WithResty 设置 pushKey
func WithPushKey(pushKey string) Option {
	return Option{F: func(o *Options) {
		o.pushKey = pushKey
	}}
}
