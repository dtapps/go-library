package domain

type TaobaoTbkDgVegasTljInstanceReportResult struct {
	/*
	   model     */
	Model *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto `json:"model,omitempty" `

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

func (s *TaobaoTbkDgVegasTljInstanceReportResult) SetModel(v TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) *TaobaoTbkDgVegasTljInstanceReportResult {
	s.Model = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportResult) SetMsgCode(v string) *TaobaoTbkDgVegasTljInstanceReportResult {
	s.MsgCode = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportResult) SetMsgInfo(v string) *TaobaoTbkDgVegasTljInstanceReportResult {
	s.MsgInfo = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportResult) SetSuccess(v bool) *TaobaoTbkDgVegasTljInstanceReportResult {
	s.Success = &v
	return s
}
