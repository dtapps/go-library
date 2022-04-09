package domain

type TaobaoTbkOrderDetailsGetOrderPage struct {
	/*
	   PublisherOrderDto     */
	Results *[]TaobaoTbkOrderDetailsGetPublisherOrderDto `json:"results,omitempty" `

	/*
	   是否还有上一页     */
	HasPre *bool `json:"has_pre,omitempty" `

	/*
	   位点字段，由调用方原样传递     */
	PositionIndex *string `json:"position_index,omitempty" `

	/*
	   是否还有下一页     */
	HasNext *bool `json:"has_next,omitempty" `

	/*
	   页码     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   页大小     */
	PageSize *int64 `json:"page_size,omitempty" `
}

func (s *TaobaoTbkOrderDetailsGetOrderPage) SetResults(v []TaobaoTbkOrderDetailsGetPublisherOrderDto) *TaobaoTbkOrderDetailsGetOrderPage {
	s.Results = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetOrderPage) SetHasPre(v bool) *TaobaoTbkOrderDetailsGetOrderPage {
	s.HasPre = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetOrderPage) SetPositionIndex(v string) *TaobaoTbkOrderDetailsGetOrderPage {
	s.PositionIndex = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetOrderPage) SetHasNext(v bool) *TaobaoTbkOrderDetailsGetOrderPage {
	s.HasNext = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetOrderPage) SetPageNo(v int64) *TaobaoTbkOrderDetailsGetOrderPage {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetOrderPage) SetPageSize(v int64) *TaobaoTbkOrderDetailsGetOrderPage {
	s.PageSize = &v
	return s
}
