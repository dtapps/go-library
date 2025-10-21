package qxwlwagnt

import (
	"context"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

type V3CardOperatorBatchResponse struct {
	Iccid  string `json:"iccid"`  // ICCID卡号
	Status string `json:"status"` // 提交状态编码
}

// V3CardOperatorBatch 批量业务状态变更
// http://docs.konyun.net/web/#/71/2398
func (c *Client) V3CardOperatorBatch(ctx context.Context, iccids []string, action string, notMustParams ...*gorequest.Params) (response CommonResponse[[]V3CardOperatorBatchResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccids", strings.Join(iccids, "_")) // 集成电路卡识别码，IC 卡的唯一识别号码(多个iccid之间_分割,最多不超过200个)
	params.Set("action", action)                    // 动作（具体参数请联系接口提供方索取）

	// 请求
	err = c.Request(ctx, "/api/v3/cardOperatorBatch", params, http.MethodPost, &response)
	return
}
