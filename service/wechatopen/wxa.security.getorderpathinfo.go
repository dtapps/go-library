package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaSecurityGetOrderPathInfoResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
	msg     struct {
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

type WxaSecurityGetOrderPathInfoResult struct {
	Result WxaSecurityGetOrderPathInfoResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
}

func newWxaSecurityGetOrderPathInfoResult(result WxaSecurityGetOrderPathInfoResponse, body []byte, http gorequest.Response) *WxaSecurityGetOrderPathInfoResult {
	return &WxaSecurityGetOrderPathInfoResult{Result: result, Body: body, Http: http}
}

// WxaSecurityGetOrderPathInfo 获取订单页 path 信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/basic-info-management/getOrderPathInfo.html
func (c *Client) WxaSecurityGetOrderPathInfo(ctx context.Context, infoType int, notMustParams ...gorequest.Params) (*WxaSecurityGetOrderPathInfoResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaSecurityGetOrderPathInfoResult(WxaSecurityGetOrderPathInfoResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("info_type", infoType)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/security/getorderpathinfo?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaSecurityGetOrderPathInfoResult(WxaSecurityGetOrderPathInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaSecurityGetOrderPathInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaSecurityGetOrderPathInfoResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSecurityGetOrderPathInfoResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 61041:
		return "订单页 path 未设置"
	}
	return "系统繁忙"
}
