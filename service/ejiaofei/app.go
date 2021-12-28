package ejiaofei

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/gomd5"
	"net/http"
)

type App struct {
	UserID  string
	Pwd     string
	Key     string
	signStr string
	ZapLog  *zap.Logger // 日志服务
}

func (app *App) request(url string, params map[string]interface{}, method string) ([]byte, error) {
	// 公共参数
	params["userid"] = app.UserID
	params["pwd"] = app.Pwd
	// 签名
	params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", app.signStr, app.Key))
	switch method {
	case http.MethodGet:
		// 请求
		get, err := gohttp.Get(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(get)
		}
		return get.Body, err
	case http.MethodPost:
		// 请求
		postJson, err := gohttp.PostForm(url, params)
		// 日志
		if app.ZapLog != nil {
			app.ZapLog.Sugar().Info(postJson)
		}
		return postJson.Body, err
	default:
		return nil, errors.New("请求类型不支持")
	}
}
