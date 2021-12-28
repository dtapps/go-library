package kashangwl

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gorm.io/gorm"
	"time"
)

// App 卡商网服务
type App struct {
	CustomerId  int           // 商家编号
	CustomerKey string        // 商家密钥
	ZapLog      *zap.Logger   // 日志服务
	Db          *gorm.DB      // 关系数据库服务
	RDb         *redis.Client // 缓存数据库服务
	MDb         *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}) ([]byte, error) {
	// 公共参数
	params["timestamp"] = time.Now().UnixNano() / 1e6
	params["customer_id"] = app.CustomerId
	// 签名参数
	params["sign"] = app.getSign(app.CustomerKey, params)
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", postJson.Body))
	}
	return postJson.Body, err
}
