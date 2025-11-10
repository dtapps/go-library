package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetAuditStatusResponse struct {
	APIResponse        // 错误
	Auditid     int    `json:"auditid"`    // 最新的审核 ID
	Status      int    `json:"status"`     // 审核状态
	Reason      string `json:"reason"`     // 当审核被拒绝时，返回的拒绝原因
	ScreenShot  string `json:"ScreenShot"` // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

// GetAuditStatus 查询审核单状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getAuditStatus.html
func (c *Client) GetAuditStatus(ctx context.Context, authorizerAccessToken string, auditid int64, notMustParams ...*gorequest.Params) (response GetAuditStatusResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("auditid", auditid)

	// 请求
	err = c.request(ctx, "wxa/get_auditstatus?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetAuditStatusErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 86000:
		return "不是由第三方代小程序进行调用"
	case 86001:
		return "不存在第三方的已经提交的代码"
	case 85012:
		return "无效的审核 id"
	default:
		return errmsg
	}
}
