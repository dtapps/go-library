package wechatqy

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gorm.io/gorm"
)

type App struct {
	Key    string
	ZapLog *zap.Logger   // 日志服务
	Db     *gorm.DB      // 关系数据库服务
	RDb    *redis.Client // 缓存数据库服务
	MDb    *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}) (body []byte, err error) {
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s %s %s", url, postJson.Header, postJson.Body))
	}
	return postJson.Body, err
}
