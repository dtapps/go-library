package domain

type TaobaoTbkDgTpwdReportGetMapData struct {
	/*
	   截止查询时刻近1小时回流pv     */
	HourPv *int64 `json:"hour_pv,omitempty" `

	/*
	   截止查询时刻近1小时回流uv     */
	HourUv *int64 `json:"hour_uv,omitempty" `

	/*
	   今日截止查询时刻累计uv     */
	Uv *int64 `json:"uv,omitempty" `

	/*
	   今日截止查询时刻累计pv     */
	Pv *int64 `json:"pv,omitempty" `
}

func (s *TaobaoTbkDgTpwdReportGetMapData) SetHourPv(v int64) *TaobaoTbkDgTpwdReportGetMapData {
	s.HourPv = &v
	return s
}
func (s *TaobaoTbkDgTpwdReportGetMapData) SetHourUv(v int64) *TaobaoTbkDgTpwdReportGetMapData {
	s.HourUv = &v
	return s
}
func (s *TaobaoTbkDgTpwdReportGetMapData) SetUv(v int64) *TaobaoTbkDgTpwdReportGetMapData {
	s.Uv = &v
	return s
}
func (s *TaobaoTbkDgTpwdReportGetMapData) SetPv(v int64) *TaobaoTbkDgTpwdReportGetMapData {
	s.Pv = &v
	return s
}
