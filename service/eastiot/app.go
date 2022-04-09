package eastiot

import (
	"errors"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"net/http"
	"time"
)

type App struct {
	AppID  string
	ApiKey string
	Mongo  gomongo.App // 日志数据库
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
		go app.mongoLog(url, params, method, get)
		return get.Body, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		go app.mongoLog(url, params, method, postJson)
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
