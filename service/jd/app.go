package jd

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	AppKey    string        // 应用Key
	SecretKey string        // 密钥
	ZapLog    *zap.Logger   // 日志服务
	Db        *gorm.DB      // 关系数据库服务
	RDb       *redis.Client // 缓存数据库服务
	MDb       *mongo.Client // 非关系数据库服务
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	get, err := gohttp.PostForm("https://api.jd.com/routerjson", params)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("https://api.jd.com/routerjson?method=%s %s %s", params["method"], get.Header, get.Body))
	}
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
