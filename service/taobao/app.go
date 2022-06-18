package taobao

import (
	"fmt"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

// App 公共请求参数
type App struct {
	appKey       string         // 应用Key
	appSecret    string         // 密钥
	adzoneId     int64          // mm_xxx_xxx_xxx的第三位
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(appKey string, appSecret string, adzoneId int64, pgsql *gorm.DB) *App {
	app := &App{appKey: appKey, appSecret: appSecret, adzoneId: adzoneId}
	app.client = gorequest.NewHttp()
	app.client.Uri = "https://eco.taobao.com/router/rest"
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "taobao"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
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

func (app *App) request(params map[string]interface{}) (resp gorequest.Response, err error) {

	// 签名
	app.Sign(params)

	// 创建请求
	client := app.client

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Get()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if app.logStatus == true {
		go app.postgresqlLog(gostring.ToString(params["method"]), request)
	}

	return request, err
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

func (app *App) GetAppKey() string {
	return app.appKey
}

func (app *App) GetAdzoneId() int64 {
	return app.adzoneId
}
