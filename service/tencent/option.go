package tencent

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	baseURL   string // 接口地址
	endpoint  string // endpoint
	secretID  string // secret_id
	secretKey string // secret_key
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

// WithRestyLog 设置 restyLog
func WithRestyLog(restyLog *resty_log.Logger) Option {
	return Option{F: func(o *Options) {
		o.restyLog = restyLog
	}}
}

// WithDebug 设置 debug
func WithDebug(debug bool) Option {
	return Option{F: func(o *Options) {
		o.debug = debug
	}}
}

// WithRWithURLesty 设置 URL
func WithURL(baseURL string) Option {
	return Option{F: func(o *Options) {
		o.baseURL = baseURL
	}}
}

// WithEndpoint 设置 endpoint
func WithEndpoint(endpoint string) Option {
	return Option{F: func(o *Options) {
		o.endpoint = endpoint
	}}
}

// WithSecretID 设置 secret_id
func WithSecretID(secretID string) Option {
	return Option{F: func(o *Options) {
		o.secretID = secretID
	}}
}

// WithSecretKey 设置 secret_key
func WithSecretKey(secretKey string) Option {
	return Option{F: func(o *Options) {
		o.secretKey = secretKey
	}}
}
