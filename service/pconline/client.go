package pconline

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

type Client struct {
	requestClient       *gorequest.App // 请求服务
	requestClientStatus bool           // 请求服务状态
	slog                struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

func NewClient() (*Client, error) {

	c := &Client{}

	return c, nil
}
