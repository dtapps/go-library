package meituan

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

// App 美团联盟
type App struct {
	Secret string        // 秘钥
	AppKey string        // 渠道标记
	ZapLog *zap.Logger   // 日志服务
	Db     *gorm.DB      // 关系数据库服务
	RDb    *redis.Client // 缓存数据库服务
	MDb    *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(get)
		}
		return get.Body, err
	case http.MethodPost:
		// 请求
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("%s %s %s", url, postJson.Header, postJson.Body))
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
