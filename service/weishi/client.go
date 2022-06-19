package weishi

import (
	"errors"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type ConfigClient struct {
	PgsqlDb *gorm.DB // pgsql数据库
}

type Client struct {
	ua        string           // 用户代理
	client    *gorequest.App   // 请求客户端
	log       *golog.ApiClient // 日志服务
	logStatus bool             // 日志状态
	config    *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"

	c.client = gorequest.NewHttp()
	if c.config.PgsqlDb != nil {
		c.logStatus = true
		c.log, err = golog.NewApiClient(&golog.ConfigApiClient{
			Db:        c.config.PgsqlDb,
			TableName: logTable,
		})
		if err != nil {
			return nil, err
		}
	}

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
