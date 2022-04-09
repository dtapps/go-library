package domain

type TaobaoTbkDgPunishOrderGetPageResult struct {
	/*
	   处罚订单列表     */
	Results *[]TaobaoTbkDgPunishOrderGetResult `json:"results,omitempty" `

	/*
	   翻页的pageno     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   翻页的pagesie     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   一共能查询出来的结果总数     */
	TotalCount *int64 `json:"total_count,omitempty" `
}

func (s *TaobaoTbkDgPunishOrderGetPageResult) SetResults(v []TaobaoTbkDgPunishOrderGetResult) *TaobaoTbkDgPunishOrderGetPageResult {
	s.Results = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetPageResult) SetPageNo(v int64) *TaobaoTbkDgPunishOrderGetPageResult {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetPageResult) SetPageSize(v int64) *TaobaoTbkDgPunishOrderGetPageResult {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetPageResult) SetTotalCount(v int64) *TaobaoTbkDgPunishOrderGetPageResult {
	s.TotalCount = &v
	return s
}
