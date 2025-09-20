package gosms

import (
	"context"

	"github.com/baidubce/bce-sdk-go/services/sms"
)

// 百度云
type Baidu struct {
	client *sms.Client // 实例
}

// 初始化
func NewBaidu(ctx context.Context, opts ...Option) (client *Baidu, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Baidu{}

	client.client, err = sms.NewClient(options.accessKeyId, options.secretAccessKey, "sms.bj.baidubce.com")

	return client, err
}
