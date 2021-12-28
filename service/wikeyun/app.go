package wikeyun

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gorm.io/gorm"
)

type App struct {
	StoreID   int
	AppKey    int
	AppSecret string
	ClientIP  string
	ZapLog    *zap.Logger   // 日志服务
	Db        *gorm.DB      // 关系数据库服务
	RDb       *redis.Client // 缓存数据库服务
	MDb       *mongo.Client // 非关系数据库服务
}

func (app *App) request(url string, params map[string]interface{}) (resp []byte, err error) {
	// 签名
	sign := app.sign(params)
	// 请求
	requestUrl := fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, app.AppKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign)
	postForm, err := gohttp.PostForm(requestUrl, params)
	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", postForm.Body))
	}
	return postForm.Body, err
}
