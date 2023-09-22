package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type TbkActivityinfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ClickUrl          string `json:"click_url"`
		WxQrcodeUrl       string `json:"wx_qrcode_url"`
		ShortClickUrl     string `json:"short_click_url"`
		TerminalType      string `json:"terminal_type"`
		MaterialOssUrl    string `json:"material_oss_url"`
		PageName          string `json:"page_name"`
		PageStartTime     string `json:"page_start_time"`
		PageEndTime       string `json:"page_end_time"`
		WxMiniprogramPath string `json:"wx_miniprogram_path"`
		Tpwd              string `json:"tpwd"`
		LongTpwd          string `json:"long_tpwd"`
	} `json:"data"`
}

type TbkActivityinfoResult struct {
	Result TbkActivityinfoResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newTbkActivityinfoResult(result TbkActivityinfoResponse, body []byte, http gorequest.Response) *TbkActivityinfoResult {
	return &TbkActivityinfoResult{Result: result, Body: body, Http: http}
}

// TbkActivityinfo 官方活动转链,饿了么/口碑活动转链
// https://www.dingdanxia.com/doc/122/173
func (c *Client) TbkActivityinfo(ctx context.Context, notMustParams ...*gorequest.Params) (*TbkActivityinfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/tbk/activityinfo", params, http.MethodPost)
	if err != nil {
		return newTbkActivityinfoResult(TbkActivityinfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkActivityinfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkActivityinfoResult(response, request.ResponseBody, request), err
}
