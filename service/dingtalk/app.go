package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Secret      string
	AccessToken string
	ZapLog      *zap.Logger   // 日志服务
	Db          *gorm.DB      // 关系数据库服务
	RDb         *redis.Client // 缓存数据库服务
	MDb         *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("%s %s", url, get.Body))
		}
		return get.Body, err
	case http.MethodPost:
		// 请求参数
		paramsMarshal, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		// 请求
		postJson, err := gohttp.PostJson(url, paramsMarshal)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("%s %s", url, postJson.Body))
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
