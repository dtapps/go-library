package pconline

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

type Client struct {
	client *gorequest.App // 请求服务
	log    struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

func NewClient() (*Client, error) {

	c := &Client{}

	c.client = gorequest.NewHttp()

	return c, nil
}
