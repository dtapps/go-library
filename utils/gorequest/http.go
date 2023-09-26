package gorequest

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	cookiemonster "github.com/MercuryEngineering/CookieMonster"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gostring"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Response 返回内容
type Response struct {
	RequestId             string                 //【请求】编号
	RequestUri            string                 //【请求】链接
	RequestParams         map[string]interface{} //【请求】参数
	RequestMethod         string                 //【请求】方式
	RequestHeader         map[string]interface{} //【请求】头部
	RequestCookie         string                 //【请求】Cookie
	RequestTime           time.Time              //【请求】时间
	ResponseHeader        http.Header            //【返回】头部
	ResponseStatus        string                 //【返回】状态
	ResponseStatusCode    int                    //【返回】状态码
	ResponseBody          []byte                 //【返回】内容
	ResponseContentLength int64                  //【返回】大小
	ResponseTime          time.Time              //【返回】时间
}

// App 实例
type App struct {
	Uri                          string                 // 全局请求地址，没有设置url才会使用
	Error                        error                  // 错误
	httpUri                      string                 // 请求地址
	httpMethod                   string                 // 请求方法
	httpHeader                   map[string]interface{} // 请求头
	httpParams                   map[string]interface{} // 请求参数
	httpCookie                   string                 // Cookie
	responseContent              Response               // 返回内容
	httpContentType              string                 // 请求内容类型
	debug                        bool                   // 是否开启调试模式
	p12Cert                      *tls.Certificate       // p12证书内容
	tlsMinVersion, tlsMaxVersion uint16                 // TLS版本
	config                       struct {
		systemOs     string // 系统类型
		systemKernel string // 系统内核
		goVersion    string // go版本
		sdkVersion   string // sdk版本
	}
}

// NewHttp 实例化
func NewHttp() *App {
	app := &App{
		httpHeader: make(map[string]interface{}),
		httpParams: make(map[string]interface{}),
	}
	app.setConfig()
	return app
}

// SetDebug 设置调试模式
func (app *App) SetDebug() {
	app.debug = true
}

// SetUri 设置请求地址
func (app *App) SetUri(uri string) {
	app.httpUri = uri
}

// SetMethod 设置请求方式
func (app *App) SetMethod(method string) {
	app.httpMethod = method
}

// SetHeader 设置请求头
func (app *App) SetHeader(key, value string) {
	app.httpHeader[key] = value
}

// SetHeaders 批量设置请求头
func (app *App) SetHeaders(headers map[string]interface{}) {
	for key, value := range headers {
		app.httpHeader[key] = value
	}
}

// SetTlsVersion 设置TLS版本
func (app *App) SetTlsVersion(minVersion, maxVersion uint16) {
	app.tlsMinVersion = minVersion
	app.tlsMaxVersion = maxVersion
}

// SetAuthToken 设置身份验证令牌
func (app *App) SetAuthToken(token string) {
	app.httpHeader["Authorization"] = fmt.Sprintf("Bearer %s", token)
}

// SetUserAgent 设置用户代理，空字符串就随机设置
func (app *App) SetUserAgent(ua string) {
	if ua == "" {
		ua = GetRandomUserAgent()
	}
	app.httpHeader["User-Agent"] = ua
}

// SetContentTypeJson 设置JSON格式
func (app *App) SetContentTypeJson() {
	app.httpContentType = httpParamsModeJson
}

// SetContentTypeForm 设置FORM格式
func (app *App) SetContentTypeForm() {
	app.httpContentType = httpParamsModeForm
}

// SetContentTypeXml 设置XML格式
func (app *App) SetContentTypeXml() {
	app.httpContentType = httpParamsModeXml
}

// SetParam 设置请求参数
func (app *App) SetParam(key string, value interface{}) {
	app.httpParams[key] = value
}

// SetParams 批量设置请求参数
func (app *App) SetParams(params map[string]interface{}) {
	for key, value := range params {
		app.httpParams[key] = value
	}
}

// SetCookie 设置Cookie
func (app *App) SetCookie(value string) {
	app.httpCookie = value
}

// SetP12Cert 设置证书
func (app *App) SetP12Cert(content *tls.Certificate) {
	app.p12Cert = content
}

// Get 发起GET请求
func (app *App) Get(ctx context.Context, uri ...string) (httpResponse Response, err error) {
	if len(uri) == 1 {
		app.Uri = uri[0]
	}
	// 设置请求方法
	app.httpMethod = http.MethodGet
	return request(app, ctx)
}

// Post 发起POST请求
func (app *App) Post(ctx context.Context, uri ...string) (httpResponse Response, err error) {
	if len(uri) == 1 {
		app.Uri = uri[0]
	}
	// 设置请求方法
	app.httpMethod = http.MethodPost
	return request(app, ctx)
}

// Request 发起请求
func (app *App) Request(ctx context.Context) (httpResponse Response, err error) {
	return request(app, ctx)
}

// 请求接口
func request(app *App, ctx context.Context) (httpResponse Response, err error) {

	// 赋值
	httpResponse.RequestTime = gotime.Current().Time
	httpResponse.RequestUri = app.httpUri
	httpResponse.RequestMethod = app.httpMethod
	httpResponse.RequestParams = app.httpParams
	httpResponse.RequestHeader = app.httpHeader
	httpResponse.RequestCookie = app.httpCookie

	// 判断网址
	if httpResponse.RequestUri == "" {
		httpResponse.RequestUri = app.Uri
	}
	if httpResponse.RequestUri == "" {
		app.Error = errors.New("没有设置Uri")
		if app.debug {
			log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
		}
		return httpResponse, app.Error
	}

	// 创建 http 客户端
	client := &http.Client{}

	transportStatus := false
	transport := &http.Transport{}
	transportTls := &tls.Config{}

	if app.p12Cert != nil {
		transportStatus = true
		// 配置
		transportTls.Certificates = []tls.Certificate{*app.p12Cert}
		transport.DisableCompression = true
	}

	if app.tlsMinVersion != 0 && app.tlsMaxVersion != 0 {
		transportStatus = true
		// 配置
		transportTls.MinVersion = app.tlsMinVersion
		transportTls.MaxVersion = app.tlsMaxVersion
	}

	if transportStatus {
		transport.TLSClientConfig = transportTls
		client = &http.Client{
			Transport: transport,
		}
	}

	// SDK版本
	httpResponse.RequestHeader["Sdk-User-Agent"] = fmt.Sprintf(userAgentFormat, app.config.systemOs, app.config.systemKernel, app.config.goVersion) + "/" + go_library.Version()

	// 请求类型
	if app.httpContentType == "" {
		app.httpContentType = httpParamsModeJson
	}
	switch app.httpContentType {
	case httpParamsModeJson:
		httpResponse.RequestHeader["Content-Type"] = "application/json"
	case httpParamsModeForm:
		httpResponse.RequestHeader["Content-Type"] = "application/x-www-form-urlencoded"
	case httpParamsModeXml:
		httpResponse.RequestHeader["Content-Type"] = "text/xml"
	}

	// 跟踪编号
	httpResponse.RequestId = gotrace_id.GetTraceIdContext(ctx)
	if httpResponse.RequestId == "" {
		httpResponse.RequestId = gostring.GetUuId()
	}
	httpResponse.RequestHeader["X-Request-Id"] = httpResponse.RequestId

	// 请求内容
	var reqBody io.Reader

	if httpResponse.RequestMethod == http.MethodPost && app.httpContentType == httpParamsModeJson {
		jsonStr, err := gojson.Marshal(httpResponse.RequestParams)
		if err != nil {
			app.Error = errors.New(fmt.Sprintf("解析出错 %s", err))
			if app.debug {
				log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
			}
			return httpResponse, app.Error
		}
		// 赋值
		reqBody = bytes.NewBuffer(jsonStr)
	}

	if httpResponse.RequestMethod == http.MethodPost && app.httpContentType == httpParamsModeForm {
		// 携带 form 参数
		form := url.Values{}
		for k, v := range httpResponse.RequestParams {
			form.Add(k, GetParamsString(v))
		}
		// 赋值
		reqBody = strings.NewReader(form.Encode())
	}

	if app.httpContentType == httpParamsModeXml {
		reqBody, err = ToXml(httpResponse.RequestParams)
		if err != nil {
			app.Error = errors.New(fmt.Sprintf("解析XML出错 %s", err))
			if app.debug {
				log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
			}
			return httpResponse, app.Error
		}
	}

	// 创建请求
	req, err := http.NewRequest(httpResponse.RequestMethod, httpResponse.RequestUri, reqBody)
	if err != nil {
		app.Error = errors.New(fmt.Sprintf("创建请求出错 %s", err))
		if app.debug {
			log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
		}
		return httpResponse, app.Error
	}

	// GET 请求携带查询参数
	if httpResponse.RequestMethod == http.MethodGet {
		q := req.URL.Query()
		for k, v := range httpResponse.RequestParams {
			q.Add(k, GetParamsString(v))
		}
		req.URL.RawQuery = q.Encode()
	}

	// 设置请求头
	if len(httpResponse.RequestHeader) > 0 {
		for key, value := range httpResponse.RequestHeader {
			req.Header.Set(key, fmt.Sprintf("%v", value))
		}
	}

	// 设置Cookie
	if httpResponse.RequestCookie != "" {
		cookies, _ := cookiemonster.ParseString(httpResponse.RequestCookie)
		if len(cookies) > 0 {
			for _, c := range cookies {
				req.AddCookie(c)
			}
		} else {
			req.Header.Set("Cookie", httpResponse.RequestCookie)
		}
	}

	if app.debug {
		log.Printf("{trace_id=%s}请求Time：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.RequestTime)
		log.Printf("{trace_id=%s}请求Uri：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.RequestUri)
		log.Printf("{trace_id=%s}请求Method：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.RequestMethod)
		log.Printf("{trace_id=%s}请求ContentType：%s\n", gotrace_id.GetTraceIdContext(ctx), app.httpContentType)
		log.Printf("{trace_id=%s}请求Params Get：%s\n", gotrace_id.GetTraceIdContext(ctx), req.URL.RawQuery)
		log.Printf("{trace_id=%s}请求Params Post：%s\n", gotrace_id.GetTraceIdContext(ctx), reqBody)
		log.Printf("{trace_id=%s}请求Header：%s\n", gotrace_id.GetTraceIdContext(ctx), req.Header)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		app.Error = errors.New(fmt.Sprintf("请求出错 %s", err))
		if app.debug {
			log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
		}
		return httpResponse, app.Error
	}

	// 最后关闭连接
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close() // nolint

	// 读取内容
	body, err := io.ReadAll(reader)
	if err != nil {
		app.Error = errors.New(fmt.Sprintf("解析内容出错 %s", err))
		if app.debug {
			log.Printf("{trace_id=%s}请求异常：%v\n", gotrace_id.GetTraceIdContext(ctx), app.Error)
		}
		return httpResponse, app.Error
	}

	// 赋值
	httpResponse.ResponseTime = gotime.Current().Time
	httpResponse.ResponseStatus = resp.Status
	httpResponse.ResponseStatusCode = resp.StatusCode
	httpResponse.ResponseHeader = resp.Header
	httpResponse.ResponseBody = body
	httpResponse.ResponseContentLength = resp.ContentLength

	if app.debug {
		log.Printf("{trace_id=%s}返回Time：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseTime)
		log.Printf("{trace_id=%s}返回Status：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseStatus)
		log.Printf("{trace_id=%s}返回StatusCode：%v\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseStatusCode)
		log.Printf("{trace_id=%s}返回Header：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseHeader)
		log.Printf("{trace_id=%s}返回Body：%s\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseBody)
		log.Printf("{trace_id=%s}返回ContentLength：%v\n\n", gotrace_id.GetTraceIdContext(ctx), httpResponse.ResponseContentLength)
	}

	return httpResponse, err
}
