package pintoto

import (
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"math"
	"net/http"
	"strconv"
	"time"
)

// App 电影票服务
type App struct {
	AppKey    string
	AppSecret string
	Mongo     gomongo.App // 日志数据库
}

// 请求
func (app *App) request(url string, params map[string]interface{}) ([]byte, error) {
	// 公共参数
	params["time"] = time.Now().Unix()
	params["appKey"] = app.AppKey
	// 签名
	params["sign"] = app.getSign(app.AppSecret, params)
	// 请求
	postForm, err := gohttp.PostForm(url, params)
	// 日志
	go app.mongoLog(url, params, http.MethodPost, postForm)
	return postForm.Body, err
}

func (app *App) GradeToFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case string:
		float, _ := strconv.ParseFloat(v, 64)
		return float
	case float64:
		return v
	case int64:
		return float64(v) / math.Pow10(0)
	default:
		return 0
	}
}
