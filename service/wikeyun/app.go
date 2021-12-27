package wikeyun

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
)

type App struct {
	StoreID   int
	AppKey    int
	AppSecret string
	ClientIP  string
}

func (app *App) request(url string, params map[string]interface{}) (resp []byte, err error) {
	// 签名
	sign := app.sign(params)
	// 请求
	requestUrl := fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, app.AppKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign)
	postForm, err := gohttp.PostForm(requestUrl, params)
	return postForm.Body, err
}
