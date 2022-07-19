package docker

import (
	"github.com/docker/docker/client"
)

type Client struct {
	Db *client.Client // 驱动
}

func NewClient() (*Client, error) {

	var err error
	c := &Client{}

	c.Db, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return c, nil
}
