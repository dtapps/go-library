package domain

type TaobaoTbkDgTpwdRiskReportResult struct {
	/*
	   x     */
	StatusList *[]TaobaoTbkDgTpwdRiskReportDataMap `json:"status_list,omitempty" `
}

func (s *TaobaoTbkDgTpwdRiskReportResult) SetStatusList(v []TaobaoTbkDgTpwdRiskReportDataMap) *TaobaoTbkDgTpwdRiskReportResult {
	s.StatusList = &v
	return s
}
