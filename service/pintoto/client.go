package pintoto

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type ConfigClient struct {
	AppKey    string
	AppSecret string
	MongoDb   *gomongo.Client // 日志数据库
	PgsqlDb   *gorm.DB        // pgsql数据库
}

// Client 电影票服务
type Client struct {
	client    *gorequest.App   // 请求客户端
	log       *golog.ApiClient // 日志服务
	logStatus bool             // 日志状态
	config    *ConfigClient    // 配置
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

func (c *Client) GradeToFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case string:
		float, _ := strconv.ParseFloat(v, 64)
		return float
	case float64:
		return v
	case int64:
		return float64(v) / math.Pow10(0)
	default:
		return 0
	}
}
