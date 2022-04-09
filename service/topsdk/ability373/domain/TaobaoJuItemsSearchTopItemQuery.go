package domain

type TaobaoJuItemsSearchTopItemQuery struct {
	/*
	   页码,必传 defalutValue:1    */
	CurrentPage *int64 `json:"current_page,omitempty" `

	/*
	   一页大小,必传 defalutValue:20    */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   媒体pid,必传     */
	Pid *string `json:"pid,omitempty" `

	/*
	   是否包邮,可不传     */
	Postage *bool `json:"postage,omitempty" `

	/*
	   状态，预热：1，正在进行中：2,可不传     */
	Status *int64 `json:"status,omitempty" `

	/*
	   淘宝类目id,可不传     */
	TaobaoCategoryId *int64 `json:"taobao_category_id,omitempty" `

	/*
	   搜索关键词,可不传     */
	Word *string `json:"word,omitempty" `
}

func (s *TaobaoJuItemsSearchTopItemQuery) SetCurrentPage(v int64) *TaobaoJuItemsSearchTopItemQuery {
	s.CurrentPage = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetPageSize(v int64) *TaobaoJuItemsSearchTopItemQuery {
	s.PageSize = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetPid(v string) *TaobaoJuItemsSearchTopItemQuery {
	s.Pid = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetPostage(v bool) *TaobaoJuItemsSearchTopItemQuery {
	s.Postage = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetStatus(v int64) *TaobaoJuItemsSearchTopItemQuery {
	s.Status = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetTaobaoCategoryId(v int64) *TaobaoJuItemsSearchTopItemQuery {
	s.TaobaoCategoryId = &v
	return s
}
func (s *TaobaoJuItemsSearchTopItemQuery) SetWord(v string) *TaobaoJuItemsSearchTopItemQuery {
	s.Word = &v
	return s
}
