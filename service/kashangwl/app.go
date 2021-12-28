package kashangwl

import (
	"encoding/json"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"gopkg.in/dtapps/go-library.v3/utils/golog"
	"time"
)

// App 卡商网服务
type App struct {
	CustomerId  int       // 商家编号
	CustomerKey string    // 商家密钥
	ZapLog      golog.App // 日志服务
}

func (app *App) request(url string, params map[string]interface{}) ([]byte, error) {
	// 公共参数
	params["timestamp"] = time.Now().UnixNano() / 1e6
	params["customer_id"] = app.CustomerId
	// 签名参数
	params["sign"] = app.getSign(app.CustomerKey, params)
	// 请求参数
	paramsStr, err := json.Marshal(params)
	// 请求
	postJson, err := gohttp.PostJson(url, paramsStr)
	// 日志
	if app.ZapLog.Logger != nil {
		app.ZapLog.LogName = "kashangwl.log"
		app.ZapLog.Logger.Sugar().Info(postJson.Body)
	}
	return postJson.Body, err
}
