package jd

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
)

type App struct {
	AppKey    string // 应用Key
	SecretKey string // 密钥
}

type ErrResp struct {
	Code          string `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	ErrorSolution string `json:"errorSolution"`
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	httpGet, err := gohttp.PostForm("https://api.jd.com/routerjson", params)
	// 检查错误
	var errResp ErrResp
	_ = json.Unmarshal(httpGet.Body, &errResp)
	return httpGet.Body, err
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
