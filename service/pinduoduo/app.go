package pinduoduo

import (
	"fmt"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
)

// App 公共请求参数
type App struct {
	clientId     string         // POP分配给应用的client_id
	clientSecret string         // POP分配给应用的client_secret
	mediaId      string         // 媒体ID
	Pid          string         // 推广位
	pgsql        *gorm.DB       // pgsql数据库
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
}

func NewApp(clientId string, clientSecret string, mediaId string, pid string, pgsql *gorm.DB) *App {
	app := &App{clientId: clientId, clientSecret: clientSecret, mediaId: mediaId, Pid: pid}
	app.client = gorequest.NewHttp()
	app.client.Uri = "https://gw-api.pinduoduo.com/api/router"
	if pgsql != nil {
		app.pgsql = pgsql
		app.logStatus = true
		app.logTableName = "pinduoduo"
		app.log = golog.NewApi(&golog.ApiConfig{
			Db:        pgsql,
			TableName: app.logTableName,
		})
	}
	return app
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string      `json:"error_msg"`
		SubMsg    string      `json:"sub_msg"`
		SubCode   interface{} `json:"sub_code"`
		ErrorCode int         `json:"error_code"`
		RequestId string      `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
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
		go app.postgresqlLog(gostring.ToString(params["type"]), request)
	}

	return request, err
}

func (app *App) SalesTipParseInt64(salesTip string) int64 {
	parseInt, err := strconv.ParseInt(salesTip, 10, 64)
	if err != nil {
		salesTipStr := salesTip
		if strings.Contains(salesTip, "万+") {
			salesTipStr = strings.Replace(salesTip, "万+", "0000", -1)
		} else if strings.Contains(salesTip, "万") {
			salesTipStr = strings.Replace(salesTip, "万", "000", -1)
		}
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(salesTipStr, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1]))
		} else if len(SalesTipMap) == 1 {
			return gostring.ToInt64(SalesTipMap[0])
		} else {
			return 0
		}
	} else {
		return parseInt
	}
}

func (app *App) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 1000
}
