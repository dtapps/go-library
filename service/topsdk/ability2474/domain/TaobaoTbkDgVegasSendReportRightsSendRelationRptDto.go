package domain

type TaobaoTbkDgVegasSendReportRightsSendRelationRptDto struct {
	/*
	   日期     */
	BizDate *string `json:"biz_date,omitempty" `

	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   红包发放数量     */
	FundNum *int64 `json:"fund_num,omitempty" `
}

func (s *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) SetBizDate(v string) *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto {
	s.BizDate = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) SetRelationId(v int64) *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) SetFundNum(v int64) *TaobaoTbkDgVegasSendReportRightsSendRelationRptDto {
	s.FundNum = &v
	return s
}
