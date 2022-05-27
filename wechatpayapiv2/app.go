package wechatpayapiv2

import (
	"crypto/tls"
	"go.dtapp.net/library/golog"
	"go.dtapp.net/library/gomongo"
	"go.dtapp.net/library/gorequest"
	"gorm.io/gorm"
)

// App 微信支付服务
type App struct {
	appId        string // 小程序或者公众号唯一凭证
	appSecret    string // 小程序或者公众号唯一凭证密钥
	mchId        string // 微信支付的商户id
	mchKey       string // 私钥
	certString   string
	keyString    string
	mongo        *gomongo.Client // 日志数据库
	pgsql        *gorm.DB        // pgsql数据库
	client       *gorequest.App  // 请求客户端
	log          *golog.Api      // 日志服务
	logTableName string          // 日志表名
	logStatus    bool            // 日志状态
}

func NewApp(appId string, appSecret string, mchId string, mchKey string, certString string, keyString string, pgsql *gorm.DB) *App {
	app := &App{appId: appId, appSecret: appSecret, mchId: mchId, mchKey: mchKey, certString: certString, keyString: keyString}
	app.client = gorequest.NewHttp()
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "wechatpayapiv2"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

func (app *App) request(url string, params map[string]interface{}, cert *tls.Certificate) (resp gorequest.Response, err error) {

	// 创建请求
	client := app.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeXml()

	// 设置参数
	client.SetParams(params)

	// 设置证书
	client.SetP12Cert(cert)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.logStatus == true {
		go app.postgresqlLog(request)
	}

	return request, err
}

func (app *App) P12ToPem() (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(app.certString), []byte(app.keyString))
	return &pemCert, err
}
