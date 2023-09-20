package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentSetPrivacySettingResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
}

type CgiBinComponentSetPrivacySettingResult struct {
	Result CgiBinComponentSetPrivacySettingResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newCgiBinComponentSetPrivacySettingResult(result CgiBinComponentSetPrivacySettingResponse, body []byte, http gorequest.Response) *CgiBinComponentSetPrivacySettingResult {
	return &CgiBinComponentSetPrivacySettingResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentSetPrivacySetting 配置小程序用户隐私保护指引
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func (c *Client) CgiBinComponentSetPrivacySetting(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*CgiBinComponentSetPrivacySettingResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/setprivacysetting?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newCgiBinComponentSetPrivacySettingResult(CgiBinComponentSetPrivacySettingResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinComponentSetPrivacySettingResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentSetPrivacySettingResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *CgiBinComponentSetPrivacySettingResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86069:
		return "owner_setting必填字段字段缺失"
	case 86070:
		return "notice_method必填字段字段缺失"
	case 86072:
		return "store_expire_timestamp参数无效。如果是编码格式不对，也会报这个错"
	case 86073:
		return "ext_file_media_id参数无效"
	case 86074:
		return "现网隐私协议不存在"
	case 86075:
		return "现网隐私协议的ext_file_media_id禁止修改"
	}
	return "系统繁忙"
}
