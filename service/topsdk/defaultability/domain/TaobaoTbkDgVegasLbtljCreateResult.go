package domain

type TaobaoTbkDgVegasLbtljCreateResult struct {
	/*
	   model     */
	Model *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult `json:"model,omitempty" `

	/*
	   msgCode     */
	MsgCode *string `json:"msg_code,omitempty" `

	/*
	   msgInfo     */
	MsgInfo *string `json:"msg_info,omitempty" `

	/*
	   是否成功     */
	Success *bool `json:"success,omitempty" `
}

func (s *TaobaoTbkDgVegasLbtljCreateResult) SetModel(v TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult) *TaobaoTbkDgVegasLbtljCreateResult {
	s.Model = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateResult) SetMsgCode(v string) *TaobaoTbkDgVegasLbtljCreateResult {
	s.MsgCode = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateResult) SetMsgInfo(v string) *TaobaoTbkDgVegasLbtljCreateResult {
	s.MsgInfo = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateResult) SetSuccess(v bool) *TaobaoTbkDgVegasLbtljCreateResult {
	s.Success = &v
	return s
}
