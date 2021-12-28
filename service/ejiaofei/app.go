package ejiaofei

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/gomd5"
	"net/http"
)

type App struct {
	UserID  string
	Pwd     string
	Key     string
	signStr string
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["userid"] = app.UserID
	params["pwd"] = app.Pwd
	// 签名
	params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", app.signStr, app.Key))
	// 请求
	if method == http.MethodGet {
		get, err := gohttp.Get(url, params)
		return get.Body, err
	} else {
		postJson, err := gohttp.PostForm(url, params)
		return postJson.Body, err
	}
}
