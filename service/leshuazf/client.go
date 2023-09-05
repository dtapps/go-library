package leshuazf

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ConfigClient struct {
	AgentId     string // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	Environment string // 环境
	KeyAgent    string
}

// Client 乐刷
type Client struct {
	client *gorequest.App // 请求服务
	config struct {
		agentId     string // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
		environment string // 环境
		keyAgent    string
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

func NewClient(config *ConfigClient) (*Client, error) {

	c := &Client{}

	c.config.agentId = config.AgentId
	c.config.environment = config.Environment
	c.config.keyAgent = config.KeyAgent

	c.client = gorequest.NewHttp()

	return c, nil
}
