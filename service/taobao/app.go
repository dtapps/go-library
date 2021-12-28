package taobao

import (
	"encoding/json"
	"fmt"
	"gopkg.in/dtapps/go-library.v2/utils/gohttp"
	"gopkg.in/dtapps/go-library.v2/utils/gostring"
	"regexp"
	"strconv"
)

// App 公共请求参数
type App struct {
	AppKey    string // 应用Key
	AppSecret string // 密钥
}
type ErrResp struct {
	ErrorResponse struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		SubCode   string `json:"sub_code"`
		SubMsg    string `json:"sub_msg"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	httpGet, err := gohttp.Get("https://eco.taobao.com/router/rest", params)
	// 检查错误
	var errResp ErrResp
	_ = json.Unmarshal(httpGet.Body, &errResp)
	return httpGet.Body, err
}

func (app *App) ZkFinalPriceParseInt64(ZkFinalPrice string) int64 {
	parseInt, err := strconv.ParseInt(ZkFinalPrice, 10, 64)
	if err != nil {
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(ZkFinalPrice, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1])) * 10
		} else {
			return gostring.ToInt64(SalesTipMap[0]) * 100
		}
	} else {
		return parseInt * 100
	}
}

func (app *App) CommissionRateParseInt64(CommissionRate string) int64 {
	parseInt, err := strconv.ParseInt(CommissionRate, 10, 64)
	if err != nil {
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(CommissionRate, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1]))
		} else {
			return gostring.ToInt64(SalesTipMap[0])
		}
	} else {
		return parseInt
	}
}

func (app *App) CouponAmountToInt64(CouponAmount int64) int64 {
	return CouponAmount * 100
}

func (app *App) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 100
}
