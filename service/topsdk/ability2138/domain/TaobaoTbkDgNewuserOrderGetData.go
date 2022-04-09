package domain

type TaobaoTbkDgNewuserOrderGetData struct {
	/*
	   resultList     */
	Results *[]TaobaoTbkDgNewuserOrderGetMapData `json:"results,omitempty" `

	/*
	   页码     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   每页大小     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   是否有下一页     */
	HasNext *bool `json:"has_next,omitempty" `
}

func (s *TaobaoTbkDgNewuserOrderGetData) SetResults(v []TaobaoTbkDgNewuserOrderGetMapData) *TaobaoTbkDgNewuserOrderGetData {
	s.Results = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetData) SetPageNo(v int64) *TaobaoTbkDgNewuserOrderGetData {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetData) SetPageSize(v int64) *TaobaoTbkDgNewuserOrderGetData {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetData) SetHasNext(v bool) *TaobaoTbkDgNewuserOrderGetData {
	s.HasNext = &v
	return s
}
