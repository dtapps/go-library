package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinAccountGetAccountBasicInfoResponse struct {
	Errcode        int    `json:"errcode"`         // 返回码
	Errmsg         string `json:"errmsg"`          // 错误信息
	Appid          string `json:"appid"`           // 帐号 appid
	AccountType    int    `json:"account_type"`    // 帐号类型（1：订阅号，2：服务号，3：小程序）
	PrincipalType  int    `json:"principal_type"`  // 主体类型
	PrincipalName  string `json:"principal_name"`  // 主体名称
	Credential     string `json:"credential"`      // 主体标识
	RealnameStatus int    `json:"realname_status"` // 实名验证状态 1=实名验证成功 2=实名验证中 3=实名验证失败
	WxVerifyInfo   struct {
		QualificationVerify   bool `json:"qualification_verify"`     // 是否资质认证，若是，拥有微信认证相关的权限
		NamingVerify          bool `json:"naming_verify"`            // 是否名称认证
		AnnualReview          bool `json:"annual_review"`            // 是否需要年审（qualification_verify == true 时才有该字段）
		AnnualReviewBeginTime int  `json:"annual_review_begin_time"` // 年审开始时间，时间戳（qualification_verify == true 时才有该字段）
		AnnualReviewEndTime   int  `json:"annual_review_end_time"`   // 年审截止时间，时间戳（qualification_verify == true 时才有该字段）
	} `json:"wx_verify_info"` // 微信认证信息
	SignatureInfo struct {
		Signature       string `json:"signature"`         // 功能介绍
		ModifyUsedCount int    `json:"modify_used_count"` // 功能介绍已使用修改次数（本月）
		ModifyQuota     int    `json:"modify_quota"`      // 功能介绍修改次数总额度（本月）
	} `json:"signature_info"` // 功能介绍信息
	HeadImageInfo struct {
		HeadImageUrl    string `json:"head_image_url"`    // 头像 url
		ModifyUsedCount int    `json:"modify_used_count"` // 头像已使用修改次数（本年）
		ModifyQuota     int    `json:"modify_quota"`      // 头像修改次数总额度（本年）
	} `json:"head_image_info"` // 头像信息
	NicknameInfo struct {
		Nickname        string `json:"nickname"`          // 小程序名称
		ModifyUsedCount int    `json:"modify_used_count"` // 小程序名称已使用修改次数（本年）
		ModifyQuota     int    `json:"modify_quota"`      // 小程序名称修改次数总额度（本年）
	} `json:"nickname_info"` // 名称信息
	RegisteredCountry int    `json:"registered_country"` // 注册国家
	Nickname          string `json:"nickname"`           // 小程序名称
}

type CgiBinAccountGetAccountBasicInfoResult struct {
	Result CgiBinAccountGetAccountBasicInfoResponse // 结果
	Body   []byte                                   // 内容
	Err    error                                    // 错误
}

func NewCgiBinAccountGetAccountBasicInfoResult(result CgiBinAccountGetAccountBasicInfoResponse, body []byte, err error) *CgiBinAccountGetAccountBasicInfoResult {
	return &CgiBinAccountGetAccountBasicInfoResult{Result: result, Body: body, Err: err}
}

// CgiBinAccountGetAccountBasicInfo 获取基本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html
func (app *App) CgiBinAccountGetAccountBasicInfo() *CgiBinAccountGetAccountBasicInfoResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=%v", app.authorizerAccessToken), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinAccountGetAccountBasicInfoResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinAccountGetAccountBasicInfoResult(response, body, err)
}
