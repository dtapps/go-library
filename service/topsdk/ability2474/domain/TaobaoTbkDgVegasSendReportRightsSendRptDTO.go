package domain

type TaobaoTbkDgVegasSendReportRightsSendRptDTO struct {
	/*
	   渠道关系id的发放数据     */
	RelationRptList *[]TaobaoTbkDgVegasSendReportRightsSendRelationRptDto `json:"relation_rpt_list,omitempty" `
}

func (s *TaobaoTbkDgVegasSendReportRightsSendRptDTO) SetRelationRptList(v []TaobaoTbkDgVegasSendReportRightsSendRelationRptDto) *TaobaoTbkDgVegasSendReportRightsSendRptDTO {
	s.RelationRptList = &v
	return s
}
