package domain

type TaobaoTbkRelationRefundPageResult struct {
	/*
	   pageNo     */
	PageNo *string `json:"page_no,omitempty" `

	/*
	   pageSize     */
	PageSize *string `json:"page_size,omitempty" `

	/*
	   总值     */
	TotalCount *string `json:"total_count,omitempty" `

	/*
	   订单列表     */
	Results *[]TaobaoTbkRelationRefundResult `json:"results,omitempty" `
}

func (s *TaobaoTbkRelationRefundPageResult) SetPageNo(v string) *TaobaoTbkRelationRefundPageResult {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkRelationRefundPageResult) SetPageSize(v string) *TaobaoTbkRelationRefundPageResult {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkRelationRefundPageResult) SetTotalCount(v string) *TaobaoTbkRelationRefundPageResult {
	s.TotalCount = &v
	return s
}
func (s *TaobaoTbkRelationRefundPageResult) SetResults(v []TaobaoTbkRelationRefundResult) *TaobaoTbkRelationRefundPageResult {
	s.Results = &v
	return s
}
