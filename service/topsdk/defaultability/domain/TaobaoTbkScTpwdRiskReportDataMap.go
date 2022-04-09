package domain

type TaobaoTbkScTpwdRiskReportDataMap struct {
	/*
	   当入参不传pid的时候返回，表示授权账号关联的pid     */
	Pid *string `json:"pid,omitempty" `

	/*
	   当入参传入pid的时候返回，表示pid关联的relationId     */
	Rid *string `json:"rid,omitempty" `

	/*
	   0表示预警，1表示拦截，如果名单中同一个淘客同时有拦截和预警信息，以拦截为准     */
	Status *string `json:"status,omitempty" `
}

func (s *TaobaoTbkScTpwdRiskReportDataMap) SetPid(v string) *TaobaoTbkScTpwdRiskReportDataMap {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkScTpwdRiskReportDataMap) SetRid(v string) *TaobaoTbkScTpwdRiskReportDataMap {
	s.Rid = &v
	return s
}
func (s *TaobaoTbkScTpwdRiskReportDataMap) SetStatus(v string) *TaobaoTbkScTpwdRiskReportDataMap {
	s.Status = &v
	return s
}
