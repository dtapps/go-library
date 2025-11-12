package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetAppealRecordsResponse struct {
	APIResponse // 错误
	Records     []struct {
		AppealRecordId    int    `json:"appeal_record_id"`   // 申诉单id
		AppealTime        int    `json:"appeal_time"`        // 申诉时间
		AppealCount       int    `json:"appeal_count"`       // 申诉次数
		AppealFrom        int    `json:"appeal_from"`        // 申诉来源（0--用户，1--服务商）
		AppealStatus      int    `json:"appeal_status"`      // 申诉状态
		AuditTime         int    `json:"audit_time"`         // 审核时间
		AuditReason       int    `json:"audit_reason"`       // 审核结果理由
		PunishDescription string `json:"punish_description"` // 处罚原因描述
		Materials         []struct {
			IllegalMaterial struct {
				Content    string `json:"content"`     // 违规内容
				ContentUrl string `json:"content_url"` // 违规链接
			} `json:"illegal_material"` // 违规材料
			AppealMaterial struct {
				Reason           string   `json:"reason"`             // 申诉理由
				ProofMaterialIds []string `json:"proof_material_ids"` // 证明材料列表(可以通过“获取临时素材”接口下载对应的材料）
			} `json:"appeal_material"` // 申诉材料（针对违规材料提供的资料）
		} `json:"materials"` // 违规材料和申诉材料
	} `json:"records"` // 申诉记录列表
}

// GetAppealRecords 获取小程序申诉记录
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/record-management/api_getappealrecords.html
func (c *Client) GetAppealRecords(ctx context.Context, illegalRecordId string, notMustParams ...*gorequest.Params) (response GetAppealRecordsResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("illegal_record_id", illegalRecordId)

	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("wxa/getappealrecords"), params, http.MethodPost, &response)
	return
}
