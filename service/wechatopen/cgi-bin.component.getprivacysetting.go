package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinComponentGetPrivacySettingResponse struct {
	Errcode     int      `json:"errcode"`      // 返回码
	Errmsg      string   `json:"errmsg"`       // 返回码信息
	CodeExist   int      `json:"code_exist"`   // 代码是否存在， 0 不存在， 1 存在 。如果最近没有通过commit接口上传代码，则会出现 code_exist=0的情况。
	PrivacyList []string `json:"privacy_list"` // 代码检测出来的用户信息类型（privacy_key）
	SettingList []struct {
		PrivacyKey   string `json:"privacy_key"`   // 用户信息类型的英文名称
		PrivacyText  string `json:"privacy_text"`  // 该用户信息类型的用途
		PrivacyLabel string `json:"privacy_label"` // 用户信息类型的中文名称
	} `json:"setting_list"` // 要收集的用户信息配置
	UpdateTime   int `json:"update_time"` // 更新时间
	OwnerSetting struct {
		ContactPhone         string `json:"contact_phone"`          // 信息收集方（开发者）的邮箱
		ContactEmail         string `json:"contact_email"`          // 信息收集方（开发者）的手机号
		ContactQq            string `json:"contact_qq"`             // 信息收集方（开发者）的qq
		ContactWeixin        string `json:"contact_weixin"`         // 信息收集方（开发者）的微信号
		NoticeMethod         string `json:"notice_method"`          // 通知方式，指的是当开发者收集信息有变动时，通过该方式通知用户
		StoreExpireTimestamp string `json:"store_expire_timestamp"` // 存储期限，指的是开发者收集用户信息存储多久
		ExtFileMediaId       string `json:"ext_file_media_id"`      // 自定义 用户隐私保护指引文件的media_id
	} `json:"owner_setting"` // 收集方（开发者）信息配置
	PrivacyDesc struct {
		PrivacyDescList []struct {
			PrivacyKey  string `json:"privacy_key"`  // 用户信息类型的英文key
			PrivacyDesc string `json:"privacy_desc"` // 用户信息类型的中文描述
		} `json:"privacy_desc_list"` // 用户信息类型
	} `json:"privacy_desc"` // 用户信息类型对应的中英文描述
	SdkPrivacyInfoList []struct {
		SdkName    string `json:"sdk_name"`     // sdk的名称
		SdkBizName string `json:"sdk_biz_name"` // sdk提供方的主体名称
		SdkList    []struct {
			PrivacyKey   string `json:"privacy_key"`  // sdk收集的信息描述
			PrivacyText  string `json:"privacy_text"` // sdk收集的信息用途说明
			PrivacyLabel string `json:"privacy_label"`
		} `json:"sdk_list"` // sdk收集的信息以及用途
	} `json:"sdk_privacy_info_list"` // sdk
}

type CgiBinComponentGetPrivacySettingResult struct {
	Result CgiBinComponentGetPrivacySettingResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newCgiBinComponentGetPrivacySettingResult(result CgiBinComponentGetPrivacySettingResponse, body []byte, http gorequest.Response) *CgiBinComponentGetPrivacySettingResult {
	return &CgiBinComponentGetPrivacySettingResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentGetPrivacySetting 查询小程序用户隐私保护指引
// @privacyVer 1表示现网版本，即，传1则该接口返回的内容是现网版本的；2表示开发版，即，传2则该接口返回的内容是开发版本的。默认是2。
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/get_privacy_setting.html
func (c *Client) CgiBinComponentGetPrivacySetting(ctx context.Context, authorizerAccessToken string, privacyVer int, notMustParams ...*gorequest.Params) (*CgiBinComponentGetPrivacySettingResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("privacy_ver", privacyVer)

	// 请求
	var response CgiBinComponentGetPrivacySettingResponse
	request, err := c.request(ctx, "cgi-bin/component/getprivacysetting?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newCgiBinComponentGetPrivacySettingResult(response, request.ResponseBody, request), err
}
