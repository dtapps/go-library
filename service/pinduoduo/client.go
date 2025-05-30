package pinduoduo

import (
	"path/filepath"

	"go.dtapp.net/library/utils/resty_extend"
	"resty.dev/v3"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ClientId         string   // POP分配给应用的client_id
	ClientSecret     string   // POP分配给应用的client_secret
	MediaId          string   // 媒体ID
	Pid              string   // 推广位
	AccessToken      string   // 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
	AccessTokenScope []string // 授权范围

	Debug       bool   // 调试
	GlcStatus   bool   // 远程日志
	LogPath     string // 日志地址
	ServiceName string // 服务名称
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
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.httpClient = resty.New().SetDebug(config.Debug)
	if config.GlcStatus {
		c.httpClient.SetLogger(&resty_extend.GlcLogger{})
	} else {
		if config.LogPath != "" {
			c.httpClient.SetLogger(resty_extend.NewLog(filepath.Join(config.LogPath), config.ServiceName))
		}
	}

	c.config.clientId = config.ClientId
	c.config.clientSecret = config.ClientSecret
	c.config.mediaId = config.MediaId
	c.config.pid = config.Pid
	c.config.accessToken = config.AccessToken
	c.config.accessTokenScope = config.AccessTokenScope

	return c, nil
}

// Close 关闭 请求客户端
func (c *Client) Close() {
	if c.httpClient != nil {
		c.httpClient.Close()
	}
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string `json:"error_msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   any    `json:"sub_code"`
		ErrorCode int    `json:"error_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}
