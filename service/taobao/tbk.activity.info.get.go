package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkActivityInfoGetResponse struct {
	TbkActivityInfoGetResponse struct {
		Data struct {
			WxQrcodeUrl       string `json:"wx_qrcode_url"`
			ClickUrl          string `json:"click_url"`
			ShortClickUrl     string `json:"short_click_url"`
			TerminalType      string `json:"terminal_type"`
			MaterialOssUrl    string `json:"material_oss_url"`
			PageName          string `json:"page_name"`
			PageStartTime     string `json:"page_start_time"`
			PageEndTime       string `json:"page_end_time"`
			WxMiniprogramPath string `json:"wx_miniprogram_path"`
		} `json:"data"`
	} `json:"tbk_activity_info_get_response"`
}

type TbkActivityInfoGetResult struct {
	Result TbkActivityInfoGetResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newTbkActivityInfoGetResult(result TbkActivityInfoGetResponse, body []byte, http gorequest.Response) *TbkActivityInfoGetResult {
	return &TbkActivityInfoGetResult{Result: result, Body: body, Http: http}
}

// TbkActivityInfoGet 淘宝客-推广者-官方活动转链
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.5a83669a7rURsF&source=search&docId=48340&docType=2
func (c *Client) TbkActivityInfoGet(ctx context.Context, notMustParams ...*gorequest.Params) (*TbkActivityInfoGetResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.activity.info.get", notMustParams...)
	params.Set("adzone_id", c.GetAdzoneId())
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkActivityInfoGetResult(TbkActivityInfoGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkActivityInfoGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkActivityInfoGetResult(response, request.ResponseBody, request), err
}
