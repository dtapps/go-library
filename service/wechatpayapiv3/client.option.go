package wechatpayapiv3

import (
	"crypto/rsa"
	"crypto/x509"

	"resty.dev/v3"
)

type Options struct {
	restyClient *resty.Client
	debug       bool

	baseURL string // 接口地址

	appId string // 小程序或者公众号唯一凭证
	mchId string // 微信支付的商户id

	apiV3 string // APIv3密钥

	certificateSerialNo string // pem 证书号

	certificate *x509.Certificate // pem 证书
	privateKey  *rsa.PrivateKey   // pem 私钥
	publicKeyID string            // 公钥ID
	publicKey   *rsa.PublicKey    // pem 公钥

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

// WithAppid 设置 appId
func WithAppid(appId string) Option {
	return Option{F: func(o *Options) {
		o.appId = appId
	}}
}

// WithMchId 设置 mchId
func WithMchId(mchId string) Option {
	return Option{F: func(o *Options) {
		o.mchId = mchId
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
