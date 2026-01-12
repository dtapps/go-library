package feishu

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.LoggerMiddleware
	debug       bool
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
