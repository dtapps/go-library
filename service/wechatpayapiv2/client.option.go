package wechatpayapiv2

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	baseURL string // 接口地址

	appId      string // 小程序或者公众号唯一凭证
	appSecret  string // 小程序或者公众号唯一凭证密钥
	mchId      string // 微信支付的商户id
	mchKey     string // 私钥
	certString string
	keyString  string
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

// WithResty 设置 URL
func WithURL(baseURL string) Option {
	return Option{F: func(o *Options) {
		o.baseURL = baseURL
	}}
}

// WithAppid 设置 appId
func WithAppid(appId string) Option {
	return Option{F: func(o *Options) {
		o.appId = appId
	}}
}

// WithAppSecret 设置 appSecret
func WithAppSecret(appSecret string) Option {
	return Option{F: func(o *Options) {
		o.appSecret = appSecret
	}}
}

// WithMchId 设置 mchId
func WithMchId(mchId string) Option {
	return Option{F: func(o *Options) {
		o.mchId = mchId
	}}
}

// WithMchKey 设置 mchKey
func WithMchKey(mchKey string) Option {
	return Option{F: func(o *Options) {
		o.mchKey = mchKey
	}}
}

// WithCertString 设置 certString
func WithCertString(certString string) Option {
	return Option{F: func(o *Options) {
		o.certString = certString
	}}
}

// WithKeyString 设置 keyString
func WithKeyString(keyString string) Option {
	return Option{F: func(o *Options) {
		o.keyString = keyString
	}}
}
