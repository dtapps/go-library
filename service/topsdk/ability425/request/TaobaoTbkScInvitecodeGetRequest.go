package request

type TaobaoTbkScInvitecodeGetRequest struct {
	/*
	   渠道关系ID     */
	RelationId *int64 `json:"relation_id,omitempty" required:"false" `
	/*
	   渠道推广的物料类型     */
	RelationApp *string `json:"relation_app" required:"true" `
	/*
	   邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请     */
	CodeType *int64 `json:"code_type" required:"true" `
}

func (s *TaobaoTbkScInvitecodeGetRequest) SetRelationId(v int64) *TaobaoTbkScInvitecodeGetRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScInvitecodeGetRequest) SetRelationApp(v string) *TaobaoTbkScInvitecodeGetRequest {
	s.RelationApp = &v
	return s
}
func (s *TaobaoTbkScInvitecodeGetRequest) SetCodeType(v int64) *TaobaoTbkScInvitecodeGetRequest {
	s.CodeType = &v
	return s
}

func (req *TaobaoTbkScInvitecodeGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.RelationApp != nil {
		paramMap["relation_app"] = *req.RelationApp
	}
	if req.CodeType != nil {
		paramMap["code_type"] = *req.CodeType
	}
	return paramMap
}

func (req *TaobaoTbkScInvitecodeGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
