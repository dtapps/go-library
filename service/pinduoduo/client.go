package pinduoduo

import (
	"github.com/dtapps/go-library/utils/godecimal"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
	"strings"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ClientId     string // POP分配给应用的client_id
	ClientSecret string // POP分配给应用的client_secret
	MediaId      string // 媒体ID
	Pid          string // 推广位
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		clientId     string // POP分配给应用的client_id
		clientSecret string // POP分配给应用的client_secret
		mediaId      string // 媒体ID
		pid          string // 推广位
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.clientId = config.ClientId
	c.config.clientSecret = config.ClientSecret
	c.config.mediaId = config.MediaId
	c.config.pid = config.Pid

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

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
