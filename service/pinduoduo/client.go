package pinduoduo

import (
	"go.dtapp.net/library/utils/godecimal"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
	"strings"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ClientId         string   // POP分配给应用的client_id
	ClientSecret     string   // POP分配给应用的client_secret
	MediaId          string   // 媒体ID
	Pid              string   // 推广位
	AccessToken      string   // 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
	AccessTokenScope []string // 授权范围
}

// Client 实例
type Client struct {
	config struct {
		clientId         string   // POP分配给应用的client_id
		clientSecret     string   // POP分配给应用的client_secret
		mediaId          string   // 媒体ID
		pid              string   // 推广位
		accessToken      string   // 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
		accessTokenScope []string // 授权范围
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.clientId = config.ClientId
	c.config.clientSecret = config.ClientSecret
	c.config.mediaId = config.MediaId
	c.config.pid = config.Pid
	c.config.accessToken = config.AccessToken
	c.config.accessTokenScope = config.AccessTokenScope

	c.trace = true
	return c, nil
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

func (c *Client) SalesTipParseInt64(salesTip string) int64 {
	if strings.Contains(salesTip, "万+") {
		return godecimal.NewString(strings.Replace(salesTip, "万+", "0000", -1)).Int64()
	} else if strings.Contains(salesTip, "万") {
		return godecimal.NewString(strings.Replace(salesTip, "万", "000", -1)).Int64()
	} else {
		return godecimal.NewString(salesTip).Int64()
	}
}
