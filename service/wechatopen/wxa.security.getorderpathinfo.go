package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaSecurityGetOrderPathInfoResponse struct {
	APIResponse // 错误
	msg         struct {
		Path        string   `json:"path"`
		ImgList     []string `json:"img_list"`
		Video       string   `json:"video"`
		TestAccount string   `json:"test_account"`
		TestPwd     string   `json:"test_pwd"`
		TestRemark  string   `json:"test_remark"`
		Status      int      `json:"status"`
		ApplyTime   int64    `json:"apply_time"`
	} `json:"msg"`
}

// WxaSecurityGetOrderPathInfo 获取订单页 path 信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/basic-info-management/getOrderPathInfo.html
func (c *Client) WxaSecurityGetOrderPathInfo(ctx context.Context, infoType int, notMustParams ...*gorequest.Params) (response WxaSecurityGetOrderPathInfoResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("info_type", infoType)

	// 请求
	err = c.request(ctx, "wxa/security/getorderpathinfo?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaSecurityGetOrderPathInfoErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 61041:
		return "订单页 path 未设置"
	default:
		return errmsg
	}
}
