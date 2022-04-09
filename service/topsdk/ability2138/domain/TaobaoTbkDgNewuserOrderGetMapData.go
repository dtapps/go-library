package domain

import (
	"topsdk/util"
)

type TaobaoTbkDgNewuserOrderGetMapData struct {
	/*
	   新注册时间，仅淘宝拉新适用     */
	RegisterTime *util.LocalTime `json:"register_time,omitempty" `

	/*
	   当前活动为淘宝拉新活动时，bind_time为新激活时间； 当前活动为支付宝拉新活动时，bind_time为绑定时间。     */
	BindTime *util.LocalTime `json:"bind_time,omitempty" `

	/*
	   首购时间，仅淘宝，天猫拉新适用     */
	BuyTime *util.LocalTime `json:"buy_time,omitempty" `

	/*
	   新人状态， 当前活动为淘宝拉新活动时，1: 新注册，2:激活，3:首购，4:确认收货； 当前活动为支付宝实名活动时，1：已绑定，2：拉新成功，3：无效用户；当前活动为支付宝新登活动时，3：手淘首购，4：手淘确认收货；当前活动为天猫拉新活动时，2:已领取，3:已首购，4:已收货     */
	Status *int64 `json:"status,omitempty" `

	/*
	   新人手机号     */
	Mobile *string `json:"mobile,omitempty" `

	/*
	   订单淘客类型:1.淘客订单；2.非淘客订单，仅淘宝，天猫拉新适用     */
	OrderTkType *int64 `json:"order_tk_type,omitempty" `

	/*
	   分享用户(unionid)，仅淘宝，天猫拉新适用     */
	UnionId *string `json:"union_id,omitempty" `

	/*
	   来源媒体ID(pid中mm_1_2_3)中第1位     */
	MemberId *int64 `json:"member_id,omitempty" `

	/*
	   来源媒体名称     */
	MemberNick *string `json:"member_nick,omitempty" `

	/*
	   来源站点ID(pid中mm_1_2_3)中第2位     */
	SiteId *int64 `json:"site_id,omitempty" `

	/*
	   来源站点名称     */
	SiteName *string `json:"site_name,omitempty" `

	/*
	   来源广告位ID(pid中mm_1_2_3)中第3位     */
	AdzoneId *int64 `json:"adzone_id,omitempty" `

	/*
	   来源广告位名称     */
	AdzoneName *string `json:"adzone_name,omitempty" `

	/*
	   淘宝订单id，仅淘宝，天猫拉新适用     */
	TbTradeParentId *int64 `json:"tb_trade_parent_id,omitempty" `

	/*
	   确认收货时间，仅天猫拉新适用     */
	AcceptTime *util.LocalTime `json:"accept_time,omitempty" `

	/*
	   领取红包时间，仅天猫拉新适用     */
	ReceiveTime *util.LocalTime `json:"receive_time,omitempty" `

	/*
	   拉新成功时间，仅支付宝拉新适用     */
	SuccessTime *util.LocalTime `json:"success_time,omitempty" `

	/*
	   活动类型，taobao-淘宝 alipay-支付宝 tmall-天猫     */
	ActivityType *string `json:"activity_type,omitempty" `

	/*
	   活动id     */
	ActivityId *string `json:"activity_id,omitempty" `

	/*
	   日期，格式为"20180202"     */
	BizDate *string `json:"biz_date,omitempty" `

	/*
	   复购订单，仅适用于手淘拉新     */
	Orders *[]TaobaoTbkDgNewuserOrderGetOrderData `json:"orders,omitempty" `

	/*
	   绑卡日期，仅适用于手淘拉新     */
	BindCardTime *util.LocalTime `json:"bind_card_time,omitempty" `

	/*
	   loginTime     */
	LoginTime *util.LocalTime `json:"login_time,omitempty" `

	/*
	   银行卡是否是绑定状态：1-绑定，0-未绑定     */
	IsCardSave *int64 `json:"is_card_save,omitempty" `

	/*
	   使用权益时间     */
	UseRightsTime *util.LocalTime `json:"use_rights_time,omitempty" `

	/*
	   领取权益时间     */
	GetRightsTime *util.LocalTime `json:"get_rights_time,omitempty" `

	/*
	   渠道关系id     */
	RelationId *string `json:"relation_id,omitempty" `
}

func (s *TaobaoTbkDgNewuserOrderGetMapData) SetRegisterTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.RegisterTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetBindTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.BindTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetBuyTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.BuyTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetStatus(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.Status = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetMobile(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.Mobile = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetOrderTkType(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.OrderTkType = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetUnionId(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.UnionId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetMemberId(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.MemberId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetMemberNick(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.MemberNick = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetSiteId(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.SiteId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetSiteName(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.SiteName = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetAdzoneId(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetAdzoneName(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.AdzoneName = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetTbTradeParentId(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.TbTradeParentId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetAcceptTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.AcceptTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetReceiveTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.ReceiveTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetSuccessTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.SuccessTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetActivityType(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.ActivityType = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetActivityId(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.ActivityId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetBizDate(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.BizDate = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetOrders(v []TaobaoTbkDgNewuserOrderGetOrderData) *TaobaoTbkDgNewuserOrderGetMapData {
	s.Orders = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetBindCardTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.BindCardTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetLoginTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.LoginTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetIsCardSave(v int64) *TaobaoTbkDgNewuserOrderGetMapData {
	s.IsCardSave = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetUseRightsTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.UseRightsTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetGetRightsTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetMapData {
	s.GetRightsTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetMapData) SetRelationId(v string) *TaobaoTbkDgNewuserOrderGetMapData {
	s.RelationId = &v
	return s
}
