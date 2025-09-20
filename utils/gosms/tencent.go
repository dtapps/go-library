package gosms

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 腾讯云
type Tencent struct {
	client *sms.Client // 实例
}

// 初始化
func NewTencent(ctx context.Context, opts ...Option) (client *Tencent, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Tencent{}

	credential := common.NewCredential(options.secretId, options.secretKey)
	cpf := profile.NewClientProfile()

	client.client, err = sms.NewClient(credential, "ap-guangzhou", cpf)

	return client, err
}
