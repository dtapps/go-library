package domain

type TaobaoTbkDgCpaActivityReportRpcResult struct {
	/*
	   分页模型     */
	Data *TaobaoTbkDgCpaActivityReportPageResult `json:"data,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityReportRpcResult) SetData(v TaobaoTbkDgCpaActivityReportPageResult) *TaobaoTbkDgCpaActivityReportRpcResult {
	s.Data = &v
	return s
}
