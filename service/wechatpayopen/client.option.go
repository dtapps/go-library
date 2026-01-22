package wechatpayopen

import (
	"crypto/rsa"
	"crypto/x509"

	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	debug       bool

	baseURL string // 接口地址

	spAppid  string // 服务商应用ID
	spMchId  string // 服务商户号
	subAppid string // 子商户应用ID
	subMchId string // 子商户号

	apiV3 string // APIv3密钥

	certificateSerialNo string            // pem 证书号
	certificate         *x509.Certificate // pem 证书
	privateKey          *rsa.PrivateKey   // pem 私钥
	publicKeyID         string            // 公钥ID
	publicKey           *rsa.PublicKey    // pem 公钥

	err error
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

// WithApiV3 设置 apiV3
func WithApiV3(apiV3 string) Option {
	return Option{F: func(o *Options) {
		o.apiV3 = apiV3
	}}
}

// WithCertificateSerialNo 设置 certificateSerialNo
func WithCertificateSerialNo(certificateSerialNo string) Option {
	return Option{F: func(o *Options) {
		o.certificateSerialNo = certificateSerialNo
	}}
}

// WithCertificate 设置 certificate
func WithCertificate(certificate string) Option {
	rsa, err := LoadCertificate(certificate)
	if err != nil {
		return Option{F: func(o *Options) {
			o.err = err
		}}
	}
	return Option{F: func(o *Options) {
		o.certificate = rsa
	}}
}

// WithPrivateKey 设置 privateKey
func WithPrivateKey(privateKey string) Option {
	rsa, err := LoadPrivateKey(privateKey)
	if err != nil {
		return Option{F: func(o *Options) {
			o.err = err
		}}
	}
	return Option{F: func(o *Options) {
		o.privateKey = rsa
	}}
}

// WithPublicKeyID 设置 publicKeyID
func WithPublicKeyID(publicKeyID string) Option {
	return Option{F: func(o *Options) {
		o.publicKeyID = publicKeyID
	}}
}

// WithPublicKey 设置 publicKey
func WithPublicKey(publicKey string) Option {
	rsa, err := LoadPublicKey(publicKey)
	if err != nil {
		return Option{F: func(o *Options) {
			o.err = err
		}}
	}
	return Option{F: func(o *Options) {
		o.publicKey = rsa
	}}
}
