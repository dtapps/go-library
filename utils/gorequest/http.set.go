package gorequest

import (
	"crypto/tls"
	"fmt"
)

// SetUri 设置请求地址
func (c *App) SetUri(uri string) {
	if uri != "" {
		c.httpUri = uri
	}
}

// SetMethod 设置请求方式
func (c *App) SetMethod(method string) {
	if method != "" {
		c.httpMethod = method
	}
}

// SetHeader 设置请求头
func (c *App) SetHeader(key, value string) {
	c.httpHeader.Set(key, value)
}

// SetHeaders 批量设置请求头
func (c *App) SetHeaders(headers Headers) {
	for key, value := range headers {
		c.httpHeader.Set(key, value)
	}
}

// SetTlsVersion 设置TLS版本
func (c *App) SetTlsVersion(minVersion, maxVersion uint16) {
	c.tlsMinVersion = minVersion
	c.tlsMaxVersion = maxVersion
}

// SetAuthToken 设置身份验证令牌
func (c *App) SetAuthToken(token string) {
	if token != "" {
		c.httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

// SetUserAgent 设置用户代理，传空字符串就随机设置
func (c *App) SetUserAgent(ua string) {
	if ua != "" {
		c.httpHeader.Set("User-Agent", ua)
	}
}

// SetContentTypeJson 设置JSON格式
func (c *App) SetContentTypeJson() {
	c.httpContentType = httpParamsModeJson
}

// SetContentTypeForm 设置FORM格式
func (c *App) SetContentTypeForm() {
	c.httpContentType = httpParamsModeForm
}

// SetContentTypeXml 设置XML格式
func (c *App) SetContentTypeXml() {
	c.httpContentType = httpParamsModeXml
}

// SetParam 设置请求参数
func (c *App) SetParam(key string, value interface{}) {
	c.httpParams.Set(key, value)
}

// SetParams 批量设置请求参数
func (c *App) SetParams(params Params) {
	for key, value := range params {
		c.httpParams.Set(key, value)
	}
}

// SetCookie 设置Cookie
func (c *App) SetCookie(cookie string) {
	if cookie != "" {
		c.httpCookie = cookie
	}
}

// SetP12Cert 设置证书
func (c *App) SetP12Cert(content *tls.Certificate) {
	c.p12Cert = content
}

// SetClientIP 设置客户端IP
func (c *App) SetClientIP(clientIP string) {
	if clientIP != "" {
		c.clientIP = clientIP
	}
}

// SetLogFunc 设置日志记录方法
func (c *App) SetLogFunc(logFunc LogFunc) {
	c.logFunc = logFunc
}
