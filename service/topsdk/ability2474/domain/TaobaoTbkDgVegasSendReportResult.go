package domain

type TaobaoTbkDgVegasSendReportResult struct {
	/*
	   是否成功     */
	Success *bool `json:"success,omitempty" `

	/*
	   model     */
	Model *TaobaoTbkDgVegasSendReportRightsSendRptDTO `json:"model,omitempty" `

	/*
	   msgInfo     */
	MsgInfo *string `json:"msg_info,omitempty" `

	/*
	   msgCode     */
	MsgCode *string `json:"msg_code,omitempty" `
}

func (s *TaobaoTbkDgVegasSendReportResult) SetSuccess(v bool) *TaobaoTbkDgVegasSendReportResult {
	s.Success = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportResult) SetModel(v TaobaoTbkDgVegasSendReportRightsSendRptDTO) *TaobaoTbkDgVegasSendReportResult {
	s.Model = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportResult) SetMsgInfo(v string) *TaobaoTbkDgVegasSendReportResult {
	s.MsgInfo = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportResult) SetMsgCode(v string) *TaobaoTbkDgVegasSendReportResult {
	s.MsgCode = &v
	return s
}
