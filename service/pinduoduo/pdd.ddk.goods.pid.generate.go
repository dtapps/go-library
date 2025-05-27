package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type GoodsPidGeneratePIdList struct {
	CreateTime int64  `json:"create_time,omitempty"` // 推广位创建时间
	PidName    string `json:"pid_name,omitempty"`    // 推广位名称
	PId        string `json:"p_id,omitempty"`        // 调用方推广位ID
	MediaId    int64  `json:"media_id,omitempty"`    // 媒体id
}
type GoodsPidGenerate struct {
	PIdGenerateResponse struct {
		PIdList        []GoodsPidGeneratePIdList `json:"p_id_list"`
		RemainPidCount int64                     `json:"remain_pid_count"` // PID剩余数量
	} `json:"p_id_generate_response"`
}

// GoodsPidGenerate 创建多多进宝推广位
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.generate
func (c *Client) GoodsPidGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (response GoodsPidGenerate, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.pid.generate", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
