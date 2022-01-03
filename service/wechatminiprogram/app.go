package wechatminiprogram

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

// App 微信小程序服务
type App struct {
	AppId       string        // 小程序唯一凭证，即 AppID
	AppSecret   string        // 小程序唯一凭证密钥，即 AppSecret
	AccessToken string        // 接口调用凭证
	JsapiTicket string        // 签名凭证
	ZapLog      *zap.Logger   // 日志服务
	Db          *gorm.DB      // 关系数据库服务
	RDb         *redis.Client // 缓存数据库服务
	MDb         *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp gohttp.Response, err error) {
	switch method {
	case http.MethodGet:
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("wechatminiprogram %s %s %s", url, get.Header, get.Body))
		}
		return get, err
	case http.MethodPost:
		// 请求参数
		paramsStr, err := json.Marshal(params)
		postJson, err := gohttp.PostJson(url, paramsStr)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(fmt.Sprintf("wechatminiprogram %s %s %s", url, postJson.Header, postJson.Body))
		}
		return postJson, err
	default:
		return resp, errors.New("请求类型不支持")
	}
}
