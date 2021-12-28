package wechatpayapiv3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/dtapps/go-library.v3/utils/gorequest"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

// App 微信支付服务
type App struct {
	AppId           string        // 小程序或者公众号唯一凭证
	AppSecret       string        // 小程序或者公众号唯一凭证密钥
	MchId           string        // 微信支付的商户id
	AesKey          string        // 私钥
	ApiV3           string        // API v3密钥
	PrivateSerialNo string        // 私钥证书号
	MchPrivateKey   string        // 商户私有证书内容 apiclient_key.pem
	ZapLog          *zap.Logger   // 日志服务
	Db              *gorm.DB      // 关系数据库服务
	RDb             *redis.Client // 缓存数据库服务
	MDb             *mongo.Client // 非关系数据库服务
}

// ErrResp 错误返回
type ErrResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  struct {
		Field    string      `json:"field,omitempty"`
		Value    interface{} `json:"value"`
		Issue    string      `json:"issue,omitempty"`
		Location string      `json:"location"`
	} `json:"detail"`
}

func (app *App) request(url string, params map[string]interface{}, method string) (resp []byte, result ErrResp, err error) {
	// 公共参数
	if method == http.MethodPost {
		if url != "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds" {
			params["appid"] = app.AppId
			params["mchid"] = app.MchId
		}
	}
	authorization, err := app.authorization(method, params, url)
	marshal, _ := json.Marshal(params)
	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewReader(marshal))
	if err != nil {
		return nil, result, err
	}
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	req.Header.Add("User-Agent", gorequest.GetRandomUserAgent())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN")

	httpClient := &http.Client{}
	var response *http.Response
	response, err = httpClient.Do(req)

	if err != nil {
		return nil, result, err
	}

	// 	处理成功
	defer response.Body.Close()
	resp, err = ioutil.ReadAll(response.Body)
	// 检查错误
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, result, err
	}

	// 日志
	if app.ZapLog != nil {
		app.ZapLog.Sugar().Info(fmt.Sprintf("%s", resp))
	}

	// 检查请求错误
	if response.StatusCode == 200 {
		return resp, result, err
	}

	return nil, result, err
}
