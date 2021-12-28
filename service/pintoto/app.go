package pintoto

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// App 电影票服务
type App struct {
	AppKey    string
	AppSecret string
	ZapLog    *zap.Logger // 日志服务
}

type ErrResp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func (app *App) request(url string, params map[string]interface{}) ([]byte, error) {
	// 公共参数
	params["time"] = time.Now().Unix()
	params["appKey"] = app.AppKey
	// 签名
	params["sign"] = app.getSign(app.AppSecret, params)
	var req *http.Request
	req, err := http.NewRequest("POST", url, strings.NewReader(app.getRequestData(params)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		return nil, nil
	}

	// 请求错误
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("请求错误:%d", response.StatusCode))
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", body))
	}

	// 检查错误
	apiErr := ErrResp{}
	err = json.Unmarshal(body, &apiErr)
	if err != nil {
		return nil, err
	}
	// 接口状态错误
	if apiErr.Code != 200 {
		return nil, errors.New(apiErr.Message)
	}
	return body, nil
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
