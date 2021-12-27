package kashangwl

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gohttp"
	"time"
)

// App 卡商网服务
type App struct {
	CustomerId  int    // 商家编号
	CustomerKey string // 商家密钥
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
	return postJson.Body, err
}
