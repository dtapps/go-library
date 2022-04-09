package domain

type TaobaoTbkDgCpaActivityReportPageResult struct {
	/*
	   数据列表     */
	Results *[]TaobaoTbkDgCpaActivityReportVegasCpaReportDTO `json:"results,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityReportPageResult) SetResults(v []TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) *TaobaoTbkDgCpaActivityReportPageResult {
	s.Results = &v
	return s
}
