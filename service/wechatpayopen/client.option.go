package wechatpayopen

import (
	"go.dtapp.net/library/contrib/resty_log"
	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	restyLog    *resty_log.Logger
	debug       bool

	baseURL        string // 接口地址
	spAppid        string // 服务商应用ID
	spMchId        string // 服务商户号
	subAppid       string // 子商户应用ID
	subMchId       string // 子商户号
	apiV2          string // APIv2密钥
	apiV3          string // APIv3密钥
	serialNo       string // 序列号
	mchSslSerialNo string // pem 证书号
	mchSslCer      string // pem 内容
	mchSslKey      string // pem key 内容
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

// WithSpAppid 设置 spAppid
func WithSpAppid(spAppid string) Option {
	return Option{F: func(o *Options) {
		o.spAppid = spAppid
	}}
}

// WithSpMchId 设置 spMchId
func WithSpMchId(spMchId string) Option {
	return Option{F: func(o *Options) {
		o.spMchId = spMchId
	}}
}

// WithSubAppid 设置 subAppid
func WithSubAppid(subAppid string) Option {
	return Option{F: func(o *Options) {
		o.subAppid = subAppid
	}}
}

// WithSubMchId 设置 subMchId
func WithSubMchId(subMchId string) Option {
	return Option{F: func(o *Options) {
		o.subMchId = subMchId
	}}
}

// WithApiV2 设置 apiV2
func WithApiV2(apiV2 string) Option {
	return Option{F: func(o *Options) {
		o.apiV2 = apiV2
	}}
}

// WithApiV3 设置 apiV3
func WithApiV3(apiV3 string) Option {
	return Option{F: func(o *Options) {
		o.apiV3 = apiV3
	}}
}

// WithSerialNo 设置 serialNo
func WithSerialNo(serialNo string) Option {
	return Option{F: func(o *Options) {
		o.serialNo = serialNo
	}}
}

// WithMchSslSerialNo 设置 mchSslSerialNo
func WithMchSslSerialNo(mchSslSerialNo string) Option {
	return Option{F: func(o *Options) {
		o.mchSslSerialNo = mchSslSerialNo
	}}
}

// WithMchSslCer 设置 mchSslCer
func WithMchSslCer(mchSslCer string) Option {
	return Option{F: func(o *Options) {
		o.mchSslCer = mchSslCer
	}}
}

// WithMchSslKey 设置 mchSslKey
func WithMchSslKey(mchSslKey string) Option {
	return Option{F: func(o *Options) {
		o.mchSslKey = mchSslKey
	}}
}
