package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetPageResponse struct {
	Errcode  int      `json:"errcode"`
	Errmsg   string   `json:"errmsg"`
	PageList []string `json:"page_list"` // page_list 页面配置列表
}

type WxaGetPageResult struct {
	Result WxaGetPageResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newWxaGetPageResult(result WxaGetPageResponse, body []byte, http gorequest.Response) *WxaGetPageResult {
	return &WxaGetPageResult{Result: result, Body: body, Http: http}
}

// WxaGetPage 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (c *Client) WxaGetPage(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetPageResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/get_page?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodGet)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetPageResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaGetPageResult(response, request.ResponseBody, request), nil
}
