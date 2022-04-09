package domain

type TaobaoTbkDgVegasTljReportInstanceDTO struct {
	/*
	   extra     */
	Extra *TaobaoTbkDgVegasTljReportExtra `json:"extra,omitempty" `
}

func (s *TaobaoTbkDgVegasTljReportInstanceDTO) SetExtra(v TaobaoTbkDgVegasTljReportExtra) *TaobaoTbkDgVegasTljReportInstanceDTO {
	s.Extra = &v
	return s
}
