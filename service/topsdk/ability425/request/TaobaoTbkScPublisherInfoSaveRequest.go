package request

type TaobaoTbkScPublisherInfoSaveRequest struct {
	/*
	   渠道备案 - 来源，取链接的来源     */
	RelationFrom *string `json:"relation_from,omitempty" required:"false" `
	/*
	   渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他     */
	OfflineScene *string `json:"offline_scene,omitempty" required:"false" `
	/*
	   渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他     */
	OnlineScene *string `json:"online_scene,omitempty" required:"false" `
	/*
	   淘宝客邀请渠道或会员的邀请码     */
	InviterCode *string `json:"inviter_code" required:"true" `
	/*
	   类型，必选 默认为1:     */
	InfoType *int64 `json:"info_type" required:"true" `
	/*
	   媒体侧渠道备注     */
	Note *string `json:"note,omitempty" required:"false" `
	/*
	   线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),省(province,必填),市(city,必填),区县街道(location,必填),详细地址(detailAddress,必填),经营类型(career,线下个人必填),店铺类型(shopType,线下店铺必填),店铺名称(shopName,线下店铺必填),店铺证书类型(shopCertifyType,线下店铺选填),店铺证书编号(certifyNumber,线下店铺选填)     */
	RegisterInfo *string `json:"register_info,omitempty" required:"false" `
}

func (s *TaobaoTbkScPublisherInfoSaveRequest) SetRelationFrom(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.RelationFrom = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetOfflineScene(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.OfflineScene = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetOnlineScene(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.OnlineScene = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetInviterCode(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.InviterCode = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetInfoType(v int64) *TaobaoTbkScPublisherInfoSaveRequest {
	s.InfoType = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetNote(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.Note = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveRequest) SetRegisterInfo(v string) *TaobaoTbkScPublisherInfoSaveRequest {
	s.RegisterInfo = &v
	return s
}

func (req *TaobaoTbkScPublisherInfoSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RelationFrom != nil {
		paramMap["relation_from"] = *req.RelationFrom
	}
	if req.OfflineScene != nil {
		paramMap["offline_scene"] = *req.OfflineScene
	}
	if req.OnlineScene != nil {
		paramMap["online_scene"] = *req.OnlineScene
	}
	if req.InviterCode != nil {
		paramMap["inviter_code"] = *req.InviterCode
	}
	if req.InfoType != nil {
		paramMap["info_type"] = *req.InfoType
	}
	if req.Note != nil {
		paramMap["note"] = *req.Note
	}
	if req.RegisterInfo != nil {
		paramMap["register_info"] = *req.RegisterInfo
	}
	return paramMap
}

func (req *TaobaoTbkScPublisherInfoSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
