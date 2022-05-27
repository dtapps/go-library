package jd

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/gostring"
	"gorm.io/gorm"
)

type App struct {
	appKey       string         // 应用Key
	secretKey    string         // 密钥
	siteId       string         // 网站ID/APP ID
	positionId   string         // 推广位id
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(appKey string, secretKey string, siteId string, positionId string, pgsql *gorm.DB) *App {
	app := &App{appKey: appKey, secretKey: secretKey, siteId: siteId, positionId: positionId}
	app.client = gorequest.NewHttp()
	app.client.Uri = "https://api.jd.com/routerjson"
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "jd"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

// 请求接口
func (app *App) request(params map[string]interface{}) (resp gorequest.Response, err error) {

	// 签名
	app.Sign(params)

	// 创建请求
	client := app.client

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.logStatus == true {
		go app.postgresqlLog(gostring.ToString(params["method"]), request)
	}

	return request, err
}

// GoodsPriceToInt64 商品券后价
func (app *App) GoodsPriceToInt64(LowestCouponPrice float64) int64 {
	return int64(LowestCouponPrice * 100)
}

// GoodsOriginalPriceToInt64 商品原价
func (app *App) GoodsOriginalPriceToInt64(Price float64) int64 {
	return int64(Price * 100)
}

// CouponProportionToInt64 佣金比率
func (app *App) CouponProportionToInt64(CommissionShare float64) int64 {
	return int64(CommissionShare * 10)
}

// CouponAmountToInt64 优惠券金额
func (app *App) CouponAmountToInt64(Commission float64) int64 {
	return int64(Commission * 100)
}

// CommissionIntegralToInt64 佣金积分
func (app *App) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 1000
}

func (app *App) GetSiteId() string {
	return app.siteId
}

func (app *App) GetPositionId() string {
	return app.positionId
}
