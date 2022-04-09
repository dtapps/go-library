package domain

type TaobaoTbkScPublisherInfoSaveData struct {
	/*
	   渠道关系ID     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   渠道昵称     */
	AccountName *string `json:"account_name,omitempty" `

	/*
	   会员运营ID     */
	SpecialId *int64 `json:"special_id,omitempty" `

	/*
	   如果重复绑定会提示：”重复绑定渠道“或”重复绑定粉丝“     */
	Desc *string `json:"desc,omitempty" `
}

func (s *TaobaoTbkScPublisherInfoSaveData) SetRelationId(v int64) *TaobaoTbkScPublisherInfoSaveData {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveData) SetAccountName(v string) *TaobaoTbkScPublisherInfoSaveData {
	s.AccountName = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveData) SetSpecialId(v int64) *TaobaoTbkScPublisherInfoSaveData {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoSaveData) SetDesc(v string) *TaobaoTbkScPublisherInfoSaveData {
	s.Desc = &v
	return s
}
