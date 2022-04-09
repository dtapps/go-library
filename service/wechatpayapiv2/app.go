package wechatpayapiv2

import (
	"crypto/tls"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"github.com/dtapps/go-library/utils/gopostgresql"
)

// App 微信支付服务
type App struct {
	AppId      string // 小程序或者公众号唯一凭证
	AppSecret  string // 小程序或者公众号唯一凭证密钥
	MchId      string // 微信支付的商户id
	MchKey     string // 私钥
	CertString string
	KeyString  string
	Pgsql      gopostgresql.App // 日志数据库
	Mongo      gomongo.App      // 日志数据库
}

func (app *App) request(url string, param Params, cert *tls.Certificate) (resp gohttp.Response, err error) {
	// 参数
	reader, err := param.MarshalXML()
	if err != nil {
		return gohttp.Response{}, err
	}
	// 请求
	postJson, err := gohttp.PostCert(url, reader, cert)
	// 日志
	go app.mongoLog(url, param, postJson)
	return postJson, err
}

func (app *App) P12ToPem() (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(app.CertString), []byte(app.KeyString))
	return &pemCert, err
}
