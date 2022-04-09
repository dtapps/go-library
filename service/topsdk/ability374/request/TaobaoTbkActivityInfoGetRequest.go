package request

type TaobaoTbkActivityInfoGetRequest struct {
	/*
	   官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取     */
	ActivityMaterialId *string `json:"activity_material_id" required:"true" `
	/*
	   mm_xxx_xxx_xxx的第三位     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   mm_xxx_xxx_xxx 仅三方分成场景使用     */
	SubPid *string `json:"sub_pid,omitempty" required:"false" `
	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" required:"false" `
	/*
	   自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道     */
	UnionId *string `json:"union_id,omitempty" required:"false" `
}

func (s *TaobaoTbkActivityInfoGetRequest) SetActivityMaterialId(v string) *TaobaoTbkActivityInfoGetRequest {
	s.ActivityMaterialId = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetRequest) SetAdzoneId(v int64) *TaobaoTbkActivityInfoGetRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetRequest) SetSubPid(v string) *TaobaoTbkActivityInfoGetRequest {
	s.SubPid = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetRequest) SetRelationId(v int64) *TaobaoTbkActivityInfoGetRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetRequest) SetUnionId(v string) *TaobaoTbkActivityInfoGetRequest {
	s.UnionId = &v
	return s
}

func (req *TaobaoTbkActivityInfoGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ActivityMaterialId != nil {
		paramMap["activity_material_id"] = *req.ActivityMaterialId
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.SubPid != nil {
		paramMap["sub_pid"] = *req.SubPid
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.UnionId != nil {
		paramMap["union_id"] = *req.UnionId
	}
	return paramMap
}

func (req *TaobaoTbkActivityInfoGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
