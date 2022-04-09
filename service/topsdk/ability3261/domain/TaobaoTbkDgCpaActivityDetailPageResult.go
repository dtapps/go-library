package domain

type TaobaoTbkDgCpaActivityDetailPageResult struct {
	/*
	   数据结果     */
	Results *[]TaobaoTbkDgCpaActivityDetailResults `json:"results,omitempty" `

	/*
	   上一页页码     */
	PrePage *int64 `json:"pre_page,omitempty" `

	/*
	   下一页页码     */
	NextPage *int64 `json:"next_page,omitempty" `

	/*
	   当前页码     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   总共页数     */
	TotalPages *int64 `json:"total_pages,omitempty" `

	/*
	   每页条数     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   是否有下一页     */
	HasNext *bool `json:"has_next,omitempty" `

	/*
	   总条数     */
	TotalCount *int64 `json:"total_count,omitempty" `

	/*
	   是否有下一页     */
	HasPre *bool `json:"has_pre,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetResults(v []TaobaoTbkDgCpaActivityDetailResults) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.Results = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetPrePage(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.PrePage = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetNextPage(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.NextPage = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetPageNo(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetTotalPages(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.TotalPages = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetPageSize(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetHasNext(v bool) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.HasNext = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetTotalCount(v int64) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.TotalCount = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailPageResult) SetHasPre(v bool) *TaobaoTbkDgCpaActivityDetailPageResult {
	s.HasPre = &v
	return s
}
