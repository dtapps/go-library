package gosms

import (
	"context"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

// 火山引擎
type Volcengine struct {
}

// 初始化
func NewVolcengine(ctx context.Context, opts ...Option) (client *Volcengine, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Volcengine{}

	sms.DefaultInstance.Client.SetAccessKey(options.accessKey)
	sms.DefaultInstance.Client.SetSecretKey(options.secretKey)

	return client, err
}
