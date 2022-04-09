package wikeyun

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gomongo"
	"net/http"
)

type App struct {
	StoreID   int
	AppKey    int
	AppSecret string
	ClientIP  string
	Mongo     gomongo.App // 日志数据库
}

func (app *App) request(url string, params map[string]interface{}) (resp []byte, err error) {
	// 签名
	sign := app.sign(params)
	// 请求
	requestUrl := fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, app.AppKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign)
	postForm, err := gohttp.PostForm(requestUrl, params)
	// 日志
	go app.mongoLog(url, params, http.MethodPost, postForm)
	return postForm.Body, err
}
