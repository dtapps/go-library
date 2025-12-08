package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetOrderPathInfoResponse struct {
	APIResponse // 错误
	Msg         struct {
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

// GetOrderPathInfo 获取订单页path信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/basic-info-management/api_getorderpathinfo.html
func (c *Client) GetOrderPathInfo(ctx context.Context, infoType int, notMustParams ...*gorequest.Params) (response GetOrderPathInfoResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("info_type", infoType)

	// 请求
	err = c.request(ctx, "wxa/security/getorderpathinfo?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetOrderPathInfoErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 61041:
		return "订单页 path 未设置"
	default:
		return errmsg
	}
}
