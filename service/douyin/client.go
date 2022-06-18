package douyin

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
	ua           string           // 用户代理
	client       *gorequest.App   // 请求客户端
	log          *golog.ApiClient // 日志服务
	logTableName string           // 日志表名
	logStatus    bool             // 日志状态
	config       *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

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

func (c *Client) urlJudge(str string) string {
	if strings.Index(str, "douyin.com") != -1 || strings.Index(str, "iesdouyin.com") != -1 {
		return str
	}
	return ""
}
