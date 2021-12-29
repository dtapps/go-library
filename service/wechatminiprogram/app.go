package wechatminiprogram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
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

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, err error) {
	// 请求参数
	marshal, _ := json.Marshal(params)
	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	var response *http.Response
	response, err = httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	// 	处理成功
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)

	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s %s", url, resp))
	}

	// 检查请求错误
	if response.StatusCode == 200 {
		return resp, err
	}

	return nil, err
}
