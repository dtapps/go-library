package eastiot

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type App struct {
	AppID  string
	ApiKey string
	ZapLog *zap.Logger   // 日志服务
	Db     *gorm.DB      // 关系数据库服务
	RDb    *redis.Client // 缓存数据库服务
	MDb    *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["appId"] = app.AppID
	params["timeStamp"] = time.Now().Unix()
	// 签名
	params["sign"] = app.getSign(app.ApiKey, params)
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("%s", get.Body))
		}
		return get.Body, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("%s", postJson.Body))
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
