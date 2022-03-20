package wechatunion

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

type App struct {
	AppId       string        // 小程序唯一凭证，即 AppID
	AppSecret   string        // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string        // 接口调用凭证
	ZapLog      *zap.Logger   // 日志服务
	Db          *gorm.DB      // 关系数据库服务
	RDb         *redis.Client // 缓存数据库服务
	MDb         *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	switch method {
	case http.MethodGet:
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("wechatunion %s %s %s", url, get.Header, get.Body))
		}
		return get.Body, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("wechatunion %s %s %s", url, postJson.Header, postJson.Body))
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
