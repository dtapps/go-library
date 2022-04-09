package kashangwl

import (
	"dtapps/dta/library/utils/gohttp"
	"dtapps/dta/library/utils/gomongo"
	"encoding/json"
	"net/http"
	"time"
)

// App 卡商网服务
type App struct {
	CustomerId  int         // 商家编号
	CustomerKey string      // 商家密钥
	Mongo       gomongo.App // 日志数据库
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
	go app.mongoLog(url, params, http.MethodPost, postJson)
	return postJson.Body, err
}
