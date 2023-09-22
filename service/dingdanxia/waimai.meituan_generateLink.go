package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanGenerateLinkResponse struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	Data     struct{} `json:"data"`
	MiniCode string   `json:"miniCode"` // 小程序码地址
}

type WaiMaiMeituanGenerateLinkResult struct {
	Result WaiMaiMeituanGenerateLinkResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newWaiMaiMeituanGenerateLinkResult(result WaiMaiMeituanGenerateLinkResponse, body []byte, http gorequest.Response) *WaiMaiMeituanGenerateLinkResult {
	return &WaiMaiMeituanGenerateLinkResult{Result: result, Body: body, Http: http}
}

// WaiMaiMeituanGenerateLink 美团外卖/闪购/酒店/优选CPS转链接口【推荐使用】
// https://www.dingdanxia.com/doc/221/173
func (c *Client) WaiMaiMeituanGenerateLink(ctx context.Context, notMustParams ...*gorequest.Params) (*WaiMaiMeituanGenerateLinkResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_generateLink", params, http.MethodPost)
	if err != nil {
		return newWaiMaiMeituanGenerateLinkResult(WaiMaiMeituanGenerateLinkResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WaiMaiMeituanGenerateLinkResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanGenerateLinkResult(response, request.ResponseBody, request), err
}
