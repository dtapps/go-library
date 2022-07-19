package drone

import (
	"context"
	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

type ConfigClient struct {
	Token string
	Host  string
}

type Client struct {
	Db     *drone.Client // 驱动
	Config *ConfigClient // 配置
}

func NewClient(config *ConfigClient) *Client {

	c := &Client{Config: config}

	cfg := new(oauth2.Config)

	client := drone.NewClient(c.Config.Host, cfg.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: c.Config.Token,
		},
	))

	c.Db = &client

	return c
}
