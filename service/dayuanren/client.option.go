package dayuanren

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	apiURL string // 接口地址
	userID int64  // 商户ID
	apiKey string // 秘钥
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

// WithResty 设置 apiURL
func WithApiURL(apiURL string) Option {
	return Option{F: func(o *Options) {
		o.apiURL = apiURL
	}}
}

// WithResty 设置 userID
func WithUserID(userID int64) Option {
	return Option{F: func(o *Options) {
		o.userID = userID
	}}
}

// WithResty 设置 apiKey
func WithApiKey(apiKey string) Option {
	return Option{F: func(o *Options) {
		o.apiKey = apiKey
	}}
}
