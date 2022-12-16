package wechatopen

import "encoding/xml"

// SignDecryptComponentVerifyTicket 验证票据
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/component_verify_ticket.html
type SignDecryptComponentVerifyTicket struct {
	XMLName               xml.Name
	AppId                 string `xml:"appId,omitempty"`                 // 第三方平台 appid
	CreateTime            int64  `xml:"CreateTime,omitempty"`            // 时间戳，单位：s
	InfoType              string `xml:"InfoType,omitempty"`              // 固定为："component_verify_ticket"
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket,omitempty"` // Ticket 内容
}

// SignDecryptAuthorizeEvent 授权变更通知推送
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/authorize_event.html
type SignDecryptAuthorizeEvent struct {
	XMLName                      xml.Name
	AppId                        string `xml:"appId,omitempty"`                        // 第三方平台 appid
	CreateTime                   int64  `xml:"CreateTime,omitempty"`                   // 时间戳，单位：s
	InfoType                     string `xml:"InfoType,omitempty"`                     // 通知类型 unauthorized=取消授权 updateauthorized=更新授权 authorized=授权成功
	AuthorizerAppid              string `xml:"AuthorizerAppid,omitempty"`              // 公众号或小程序的 appid
	AuthorizationCode            string `xml:"AuthorizationCode,omitempty"`            // 授权码，可用于获取授权信息
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime,omitempty"` // 授权码过期时间 单位秒
	PreAuthCode                  string `xml:"PreAuthCode,omitempty"`                  // 预授权码
}

// SignDecryptNotifyThirdFasteRegister 快速注册企业/个人小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/register-management/fast-registration-ind/fastRegisterPersonalMp.html
type SignDecryptNotifyThirdFasteRegister struct {
	XMLName    xml.Name
	AppId      string `xml:"AppId,omitempty"`      // 第三方平台 appid
	CreateTime int64  `xml:"CreateTime,omitempty"` // 时间戳，单位：s
	InfoType   string `xml:"InfoType,omitempty"`   // 类型
	Appid      string `xml:"appid,omitempty"`      // 创建小程序appid
	Status     int64  `xml:"status,omitempty"`     // 状态
	AuthCode   string `xml:"auth_code,omitempty"`  // 第三方授权码
	Msg        string `xml:"msg,omitempty"`        // 信息
	Info       struct {
		Name               string `xml:"name,omitempty"` // 企业名称
		Code               string `xml:"code,omitempty"` // 企业代码
		CodeType           string `xml:"code_type,omitempty"`
		LegalPersonaWechat string `xml:"legal_persona_wechat,omitempty"` // (企业)法人微信号
		LegalPersonaName   string `xml:"legal_persona_name,omitempty"`   // (企业)法人姓名
		ComponentPhone     string `xml:"component_phone,omitempty"`      // (企业/个人)第三方联系电话
		Wxuser             string `xml:"wxuser"`                         // (个人)用户微信号
		Idname             string `xml:"idname"`                         // (个人)用户姓名
	} `xml:"info,omitempty"`
}

// SignDecryptNotifyThirdFastRegisterBetaApp 注册试用小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/register-management/fast-regist-beta/registerBetaMiniprogram.html
type SignDecryptNotifyThirdFastRegisterBetaApp struct {
	XMLName    xml.Name
	AppId      string `xml:"AppId,omitempty"`      // 第三方平台 appid
	CreateTime int64  `xml:"CreateTime,omitempty"` // 时间戳，单位：s
	InfoType   string `xml:"InfoType,omitempty"`   // 类型
	Appid      string `xml:"appid,omitempty"`      // 创建小程序appid
	Status     int64  `xml:"status,omitempty"`     // 状态
	Msg        string `xml:"msg,omitempty"`        // 信息
	Info       struct {
		UniqueId string `xml:"unique_id,omitempty"`
		Name     string `xml:"name,omitempty"` // 小程序名称
	} `xml:"info,omitempty"`
}

// SignDecryptNotifyThirdFastVerifyBetaApp 试用小程序快速认证
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/register-management/fast-regist-beta/verfifyBetaMiniprogram.html
type SignDecryptNotifyThirdFastVerifyBetaApp struct {
	XMLName    xml.Name
	AppId      string `xml:"AppId,omitempty"`      // 第三方平台 appid
	CreateTime int64  `xml:"CreateTime,omitempty"` // 时间戳，单位：s
	InfoType   string `xml:"InfoType,omitempty"`   // 类型
	Appid      string `xml:"appid,omitempty"`      // 创建小程序appid
	Status     int64  `xml:"status,omitempty"`     // 状态
	Msg        string `xml:"msg,omitempty"`        // 信息
	Info       struct {
		Name               string `xml:"name,omitempty"` // 企业名称
		Code               string `xml:"code,omitempty"` // 企业代码
		CodeType           string `xml:"code_type,omitempty"`
		LegalPersonaWechat string `xml:"legal_persona_wechat,omitempty"` // 法人微信号
		LegalPersonaName   string `xml:"legal_persona_name,omitempty"`   // 法人姓名
		ComponentPhone     string `xml:"component_phone,omitempty"`      // 第三方联系电话
	} `xml:"info,omitempty"`
}

// SignDecryptWxaNicknameAudit 名称审核结果事件推送
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/basic-info-management/setNickName.html
type SignDecryptWxaNicknameAudit struct {
	XMLName      xml.Name
	ToUserName   string `xml:"ToUserName,omitempty"`   // 小程序的原始 ID
	FromUserName string `xml:"FromUserName,omitempty"` // 发送方帐号（一个 OpenID，此时发送方是系统帐号）
	CreateTime   int64  `xml:"CreateTime,omitempty"`   // 消息创建时间 （整型），时间戳
	MsgType      string `xml:"MsgType,omitempty"`      // 消息类型 event
	Event        string `xml:"Event,omitempty"`        // 事件类型
	Ret          string `xml:"ret,omitempty"`          // 审核结果 2：失败，3：成功
	Nickname     string `xml:"nickname,omitempty"`     // 需要更改的昵称
	Reason       string `xml:"reason,omitempty"`       // 审核失败的驳回原因
}

// SignDecryptWeAppAudit 代码审核结果推送
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/submitAudit.html
type SignDecryptWeAppAudit struct {
	XMLName      xml.Name
	ToUserName   string `xml:"ToUserName,omitempty"`   // 小程序的原始 ID
	FromUserName string `xml:"FromUserName,omitempty"` // 发送方帐号（一个 OpenID，此时发送方是系统帐号）
	CreateTime   int64  `xml:"CreateTime,omitempty"`   // 消息创建时间 （整型），时间戳
	MsgType      string `xml:"MsgType,omitempty"`      // 消息类型 event
	Event        string `xml:"Event,omitempty"`        // 事件类型 weapp_audit_success=审核通过 weapp_audit_fail=审核不通过 weapp_audit_delay=审核延后
	SuccTime     int64  `xml:"SuccTime,omitempty"`     // 审核成功时的时间戳
	FailTime     int64  `xml:"FailTime,omitempty"`     // 审核不通过的时间戳
	DelayTime    int64  `xml:"DelayTime,omitempty"`    // 审核延后时的时间戳
	Reason       string `xml:"reason,omitempty"`       // 审核不通过的原因
	ScreenShot   string `xml:"ScreenShot,omitempty"`   // 审核不通过的截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

// SignDecryptWxaSecurityApplySetOrderPathInfo 申请设置订单页 path 信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/basic-info-management/applySetOrderPathInfo.html
type SignDecryptWxaSecurityApplySetOrderPathInfo struct {
	XMLName   xml.Name
	List      string `xml:"List,omitempty"`      // 申请结果列表
	Appid     string `xml:"Appid,omitempty"`     // 申请的appid
	AuditId   string `xml:"AuditId,omitempty"`   // 审核单id
	Status    string `xml:"Status,omitempty"`    // 订单页 path 状态
	ApplyTime int64  `xml:"ApplyTime,omitempty"` // 申请时间
	AuditTime int64  `xml:"AuditTime,omitempty"` // 审核时间
	Reason    string `xml:"Reason,omitempty"`    // 审核原因
}
