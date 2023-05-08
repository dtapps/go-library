package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetLatestAuditStatusResponse struct {
	Errcode         int    `json:"errcode"`           // 返回码
	Errmsg          string `json:"errmsg"`            // 错误信息
	Auditid         int    `json:"auditid"`           // 最新的审核 ID
	Status          int    `json:"status"`            // 审核状态
	Reason          string `json:"reason"`            // 当审核被拒绝时，返回的拒绝原因
	ScreenShot      string `json:"ScreenShot"`        // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
	UserDesc        string `json:"user_desc"`         // 审核版本
	UserVersion     string `json:"user_version"`      // 版本描述
	SubmitAuditTime int64  `json:"submit_audit_time"` // 时间戳，提交审核的时间
}

type WxaGetLatestAuditStatusResult struct {
	Result WxaGetLatestAuditStatusResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newWxaGetLatestAuditStatusResult(result WxaGetLatestAuditStatusResponse, body []byte, http gorequest.Response) *WxaGetLatestAuditStatusResult {
	return &WxaGetLatestAuditStatusResult{Result: result, Body: body, Http: http}
}

// WxaGetLatestAuditStatus 查询最新一次提交的审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html
func (c *Client) WxaGetLatestAuditStatus(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetLatestAuditStatusResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaGetLatestAuditStatusResult(WxaGetLatestAuditStatusResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/get_latest_auditstatus?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodGet)
	if err != nil {
		return newWxaGetLatestAuditStatusResult(WxaGetLatestAuditStatusResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGetLatestAuditStatusResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGetLatestAuditStatusResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaGetLatestAuditStatusResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86000:
		return "不是由第三方代小程序进行调用"
	case 86001:
		return "不存在第三方的已经提交的代码"
	case 85012:
		return "无效的审核 id"
	}
	return "系统繁忙"
}
