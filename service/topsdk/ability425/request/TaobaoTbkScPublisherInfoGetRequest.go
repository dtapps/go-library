package request

type TaobaoTbkScPublisherInfoGetRequest struct {
	/*
	   渠道独占 - 渠道关系ID     */
	RelationId *int64 `json:"relation_id,omitempty" required:"false" `
	/*
	   第几页 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页大小 defalutValue��10    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   类型，必选 1:渠道信息；2:会员信息     */
	InfoType *int64 `json:"info_type" required:"true" `
	/*
	   备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可     */
	RelationApp *string `json:"relation_app" required:"true" `
	/*
	   会员独占 - 会员运营ID     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   淘宝客外部用户标记，如自身系统账户ID；微信ID等     */
	ExternalId *string `json:"external_id,omitempty" required:"false" `
	/*
	   1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0 defalutValue��0    */
	ExternalType *int64 `json:"external_type,omitempty" required:"false" `
}

func (s *TaobaoTbkScPublisherInfoGetRequest) SetRelationId(v int64) *TaobaoTbkScPublisherInfoGetRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetPageNo(v int64) *TaobaoTbkScPublisherInfoGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetPageSize(v int64) *TaobaoTbkScPublisherInfoGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetInfoType(v int64) *TaobaoTbkScPublisherInfoGetRequest {
	s.InfoType = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetRelationApp(v string) *TaobaoTbkScPublisherInfoGetRequest {
	s.RelationApp = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetSpecialId(v string) *TaobaoTbkScPublisherInfoGetRequest {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetExternalId(v string) *TaobaoTbkScPublisherInfoGetRequest {
	s.ExternalId = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRequest) SetExternalType(v int64) *TaobaoTbkScPublisherInfoGetRequest {
	s.ExternalType = &v
	return s
}

func (req *TaobaoTbkScPublisherInfoGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.InfoType != nil {
		paramMap["info_type"] = *req.InfoType
	}
	if req.RelationApp != nil {
		paramMap["relation_app"] = *req.RelationApp
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.ExternalId != nil {
		paramMap["external_id"] = *req.ExternalId
	}
	if req.ExternalType != nil {
		paramMap["external_type"] = *req.ExternalType
	}
	return paramMap
}

func (req *TaobaoTbkScPublisherInfoGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
