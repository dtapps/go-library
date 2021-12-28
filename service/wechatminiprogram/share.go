package wechatminiprogram

import (
	"crypto/sha1"
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/gorandom"
	"io"
	"time"
)

type ShareResult struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	RawString string `json:"raw_string"`
	Signature string `json:"signature"`
}

func (app *App) Share(url string) (result ShareResult) {
	result.AppId = app.AppId
	result.NonceStr = gorandom.Alphanumeric(32)
	result.Timestamp = time.Now().Unix()
	result.Url = url
	result.RawString = fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%v&url=%v", app.JsapiTicket, result.NonceStr, result.Timestamp, result.Url)
	t := sha1.New()
	io.WriteString(t, result.RawString)
	result.Signature = fmt.Sprintf("%x", t.Sum(nil))
	return result
}
