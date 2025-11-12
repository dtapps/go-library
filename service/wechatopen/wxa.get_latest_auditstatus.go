package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetLatestAuditStatusResponse struct {
	APIResponse            // 错误
	Auditid         int    `json:"auditid"`           // 最新的审核 ID
	Status          int    `json:"status"`            // 审核状态
	Reason          string `json:"reason"`            // 当审核被拒绝时，返回的拒绝原因
	ScreenShot      string `json:"ScreenShot"`        // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
	UserDesc        string `json:"user_desc"`         // 审核版本
	UserVersion     string `json:"user_version"`      // 版本描述
	SubmitAuditTime int64  `json:"submit_audit_time"` // 时间戳，提交审核的时间
}

// GetLatestAuditStatus 查询最新一次审核单状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/code-management/api_getlatestauditstatus.html
func (c *Client) GetLatestAuditStatus(ctx context.Context, notMustParams ...*gorequest.Params) (response GetLatestAuditStatusResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_latest_auditstatus?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetLatestAuditStatusErrcodeInfo(errcode int, errmsg string) string {
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
