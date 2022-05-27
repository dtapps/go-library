package jd

import (
	"fmt"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"net/http"
)

type App struct {
	AppKey     string      // 应用Key
	SecretKey  string      // 密钥
	SiteId     string      // 网站ID/APP ID
	PositionId string      // 推广位id
	Mongo      gomongo.App // 日志数据库
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	get, err := gohttp.PostForm("https://api.jd.com/routerjson", params)
	// 日志
	go app.mongoLog(fmt.Sprintf("https://api.jd.com/routerjson?method=%s", params["method"]), params, http.MethodPost, get)
	return get.Body, err
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
