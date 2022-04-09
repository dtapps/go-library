package domain

type TaobaoTbkDgTpwdRiskReportDataMap struct {
	/*
	   当入参不传pid的时候返回，表示账号关联的pid     */
	Pid *string `json:"pid,omitempty" `

	/*
	   当入参传入pid的时候返回，表示pid关联的relationId     */
	Rid *string `json:"rid,omitempty" `

	/*
	   0表示预警，1表示拦截，如果名单中同一个淘客同时有拦截和预警信息，以拦截为准     */
	Status *string `json:"status,omitempty" `
}

func (s *TaobaoTbkDgTpwdRiskReportDataMap) SetPid(v string) *TaobaoTbkDgTpwdRiskReportDataMap {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkDgTpwdRiskReportDataMap) SetRid(v string) *TaobaoTbkDgTpwdRiskReportDataMap {
	s.Rid = &v
	return s
}
func (s *TaobaoTbkDgTpwdRiskReportDataMap) SetStatus(v string) *TaobaoTbkDgTpwdRiskReportDataMap {
	s.Status = &v
	return s
}
