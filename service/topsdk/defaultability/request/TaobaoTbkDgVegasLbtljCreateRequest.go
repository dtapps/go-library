package request

import (
	"topsdk/util"
)

type TaobaoTbkDgVegasLbtljCreateRequest struct {
	/*
	   妈妈广告位Id     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   安全等级，0：适用于常规淘礼金投放场景；1：更高安全级别，适用于淘礼金面额偏大等安全性较高的淘礼金投放场景，可能导致更多用户拦截。security_switch为true，此字段不填写时，使用0作为默认安全级别。如果security_switch为false，不进行安全判断。 defalutValue��0    */
	SecurityLevel *int64 `json:"security_level,omitempty" required:"false" `
	/*
	   使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始     */
	UseStartTime *string `json:"use_start_time,omitempty" required:"false" `
	/*
	   结束日期的模式,1:相对时间，2:绝对时间     */
	UseEndTimeMode *int64 `json:"use_end_time_mode" required:"true" `
	/*
	   使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束     */
	UseEndTime *string `json:"use_end_time" required:"true" `
	/*
	   裂变任务领取截止时间     */
	AcceptEndTime *util.LocalTime `json:"accept_end_time" required:"true" `
	/*
	   裂变任务领取开始时间     */
	AcceptStartTime *util.LocalTime `json:"accept_start_time" required:"true" `
	/*
	   单个淘礼金面额，支持两位小数，单位元     */
	RightsPerFace *string `json:"rights_per_face" required:"true" `
	/*
	   安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择 defalutValue��true    */
	SecuritySwitch *bool `json:"security_switch,omitempty" required:"false" `
	/*
	   单用户累计中奖次数上限 defalutValue��1    */
	UserTotalWinNumLimit *int64 `json:"user_total_win_num_limit,omitempty" required:"false" `
	/*
	   淘礼金名称，最大10个字符     */
	Name *string `json:"name,omitempty" required:"false" `
	/*
	   淘礼金总个数     */
	RightsNum *int64 `json:"rights_num" required:"true" `
	/*
	   宝贝ID     */
	ItemId *int64 `json:"item_id" required:"true" `
	/*
	   CPS佣金类型 defalutValue��MKT    */
	CampaignType *string `json:"campaign_type,omitempty" required:"false" `
	/*
	   裂变淘礼金总个数     */
	TaskRightsNum *int64 `json:"task_rights_num" required:"true" `
	/*
	   裂变单个淘礼金面额，支持两位小数，单位元     */
	TaskRightsPerFace *string `json:"task_rights_per_face" required:"true" `
	/*
	   裂变淘礼金邀请人数     */
	InviteNum *int64 `json:"invite_num" required:"true" `
	/*
	   裂变淘礼金邀请时长；单位：分钟；最大120分钟     */
	InviteTimeLimit *int64 `json:"invite_time_limit" required:"true" `
}

func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetAdzoneId(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetSecurityLevel(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.SecurityLevel = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetUseStartTime(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.UseStartTime = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetUseEndTimeMode(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.UseEndTimeMode = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetUseEndTime(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.UseEndTime = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetAcceptEndTime(v util.LocalTime) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.AcceptEndTime = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetAcceptStartTime(v util.LocalTime) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.AcceptStartTime = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetRightsPerFace(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.RightsPerFace = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetSecuritySwitch(v bool) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.SecuritySwitch = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetUserTotalWinNumLimit(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.UserTotalWinNumLimit = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetName(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.Name = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetRightsNum(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.RightsNum = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetItemId(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetCampaignType(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.CampaignType = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetTaskRightsNum(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.TaskRightsNum = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetTaskRightsPerFace(v string) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.TaskRightsPerFace = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetInviteNum(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.InviteNum = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateRequest) SetInviteTimeLimit(v int64) *TaobaoTbkDgVegasLbtljCreateRequest {
	s.InviteTimeLimit = &v
	return s
}

func (req *TaobaoTbkDgVegasLbtljCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.SecurityLevel != nil {
		paramMap["security_level"] = *req.SecurityLevel
	}
	if req.UseStartTime != nil {
		paramMap["use_start_time"] = *req.UseStartTime
	}
	if req.UseEndTimeMode != nil {
		paramMap["use_end_time_mode"] = *req.UseEndTimeMode
	}
	if req.UseEndTime != nil {
		paramMap["use_end_time"] = *req.UseEndTime
	}
	if req.AcceptEndTime != nil {
		paramMap["accept_end_time"] = *req.AcceptEndTime
	}
	if req.AcceptStartTime != nil {
		paramMap["accept_start_time"] = *req.AcceptStartTime
	}
	if req.RightsPerFace != nil {
		paramMap["rights_per_face"] = *req.RightsPerFace
	}
	if req.SecuritySwitch != nil {
		paramMap["security_switch"] = *req.SecuritySwitch
	}
	if req.UserTotalWinNumLimit != nil {
		paramMap["user_total_win_num_limit"] = *req.UserTotalWinNumLimit
	}
	if req.Name != nil {
		paramMap["name"] = *req.Name
	}
	if req.RightsNum != nil {
		paramMap["rights_num"] = *req.RightsNum
	}
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.CampaignType != nil {
		paramMap["campaign_type"] = *req.CampaignType
	}
	if req.TaskRightsNum != nil {
		paramMap["task_rights_num"] = *req.TaskRightsNum
	}
	if req.TaskRightsPerFace != nil {
		paramMap["task_rights_per_face"] = *req.TaskRightsPerFace
	}
	if req.InviteNum != nil {
		paramMap["invite_num"] = *req.InviteNum
	}
	if req.InviteTimeLimit != nil {
		paramMap["invite_time_limit"] = *req.InviteTimeLimit
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasLbtljCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
