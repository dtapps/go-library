package domain

type TaobaoTbkScTpwdRiskReportResult struct {
	/*
	   x     */
	StatusList *[]TaobaoTbkScTpwdRiskReportDataMap `json:"status_list,omitempty" `
}

func (s *TaobaoTbkScTpwdRiskReportResult) SetStatusList(v []TaobaoTbkScTpwdRiskReportDataMap) *TaobaoTbkScTpwdRiskReportResult {
	s.StatusList = &v
	return s
}
