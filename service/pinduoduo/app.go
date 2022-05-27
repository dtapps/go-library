package pinduoduo

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gohttp"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gostring"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// App 公共请求参数
type App struct {
	ClientId     string      // POP分配给应用的client_id
	ClientSecret string      // POP分配给应用的client_secret
	MediaId      string      // 媒体ID
	Pid          string      // 推广位
	Mongo        gomongo.App // 日志数据库
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string      `json:"error_msg"`
		SubMsg    string      `json:"sub_msg"`
		SubCode   interface{} `json:"sub_code"`
		ErrorCode int         `json:"error_code"`
		RequestId string      `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	get, err := gohttp.Get("https://gw-api.pinduoduo.com/api/router", params)
	// 日志
	go app.mongoLog(fmt.Sprintf("https://gw-api.pinduoduo.com/api/router?type=%s", params["type"]), params, http.MethodPost, get)
	// 检查错误
	var errResp ErrResp
	_ = json.Unmarshal(get.Body, &errResp)
	return get.Body, err
}

func (app *App) SalesTipParseInt64(salesTip string) int64 {
	parseInt, err := strconv.ParseInt(salesTip, 10, 64)
	if err != nil {
		salesTipStr := salesTip
		if strings.Contains(salesTip, "万+") {
			salesTipStr = strings.Replace(salesTip, "万+", "0000", -1)
		} else if strings.Contains(salesTip, "万") {
			salesTipStr = strings.Replace(salesTip, "万", "000", -1)
		}
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(salesTipStr, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1]))
		} else if len(SalesTipMap) == 1 {
			return gostring.ToInt64(SalesTipMap[0])
		} else {
			return 0
		}
	} else {
		return parseInt
	}
}

func (app *App) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 1000
}
