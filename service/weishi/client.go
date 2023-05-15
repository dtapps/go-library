package weishi

import (
	"errors"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"strings"
)

// ClientConfig 实例配置
type ClientConfig struct {
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		ua string // 用户代理
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"

	c.requestClient = gorequest.NewHttp()

	return c, nil
}

func (c *Client) urlJudge(str string) string {
	if strings.Index(str, "weishi.qq.com") != -1 {
		return str
	}
	return ""
}

func (c *Client) request302(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}

	response, err := client.Do(req)
	if err != nil {
		if response.StatusCode == http.StatusFound {
			location, err := response.Location()
			return location.String(), err
		} else {
			return "", err
		}
	}

	return "", nil
}
