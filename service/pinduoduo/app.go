package pinduoduo

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/gostring"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
)

// App 公共请求参数
type App struct {
	ClientId     string        // POP分配给应用的client_id
	ClientSecret string        // POP分配给应用的client_secret
	Pid          string        // 推广位
	ZapLog       *zap.Logger   // 日志服务
	Db           *gorm.DB      // 关系数据库服务
	RDb          *redis.Client // 缓存数据库服务
	MDb          *mongo.Client // 非关系数据库服务
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

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	get, err := gohttp.Get("https://gw-api.pinduoduo.com/api/router", params)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", get.Body))
	}
	// 检查错误
	var errResp ErrResp
	_ = json.Unmarshal(get.Body, &errResp)
	return get.Body, err
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
