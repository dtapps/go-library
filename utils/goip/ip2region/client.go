package ip2region

import (
	"fmt"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

// Client 客户端
type Client struct {
	v4Filepath string
	v6Filepath string
	service    *service.Ip2Region
}

// New 创建
func New(v4Filepath string, v6Filepath string) (*Client, error) {

	var err error
	c := &Client{
		v4Filepath: v4Filepath,
		v6Filepath: v6Filepath,
	}

	if c.v4Filepath == "" && c.v6Filepath == "" {
		return nil, fmt.Errorf("v4Filepath and v6Filepath are empty")
	}

	c.service, err = service.NewIp2RegionWithPath(c.v4Filepath, c.v6Filepath)
	if err != nil {
		return nil, err
	}

	return c, err
}

// Close 关闭
func (c *Client) Close() {
	if c.service != nil {
		c.service.Close()
	}
}
