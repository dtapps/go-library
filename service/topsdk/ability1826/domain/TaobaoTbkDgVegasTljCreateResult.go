package domain

type TaobaoTbkDgVegasTljCreateResult struct {
	/*
	   model     */
	Model *TaobaoTbkDgVegasTljCreateRightsInstanceCreateResult `json:"model,omitempty" `

	/*
	   msgCode     */
	MsgCode *string `json:"msg_code,omitempty" `

	/*
	   msgInfo     */
	MsgInfo *string `json:"msg_info,omitempty" `

	/*
	   success     */
	Success *bool `json:"success,omitempty" `
}

func (s *TaobaoTbkDgVegasTljCreateResult) SetModel(v TaobaoTbkDgVegasTljCreateRightsInstanceCreateResult) *TaobaoTbkDgVegasTljCreateResult {
	s.Model = &v
	return s
}
func (s *TaobaoTbkDgVegasTljCreateResult) SetMsgCode(v string) *TaobaoTbkDgVegasTljCreateResult {
	s.MsgCode = &v
	return s
}
func (s *TaobaoTbkDgVegasTljCreateResult) SetMsgInfo(v string) *TaobaoTbkDgVegasTljCreateResult {
	s.MsgInfo = &v
	return s
}
func (s *TaobaoTbkDgVegasTljCreateResult) SetSuccess(v bool) *TaobaoTbkDgVegasTljCreateResult {
	s.Success = &v
	return s
}
