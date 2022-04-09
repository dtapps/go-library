package domain

type TaobaoJuItemsSearchPaginationResult struct {
	/*
	   当前页码     */
	CurrentPage *int64 `json:"current_page,omitempty" `

	/*
	   扩展属性     */
	Extend *TaobaoJuItemsSearchExtend `json:"extend,omitempty" `

	/*
	   商品数据     */
	ModelList *[]TaobaoJuItemsSearchItems `json:"model_list,omitempty" `

	/*
	   错误码     */
	MsgCode *string `json:"msg_code,omitempty" `

	/*
	   错误信息     */
	MsgInfo *string `json:"msg_info,omitempty" `

	/*
	   一页大小     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   请求是否成功     */
	Success *bool `json:"success,omitempty" `

	/*
	   商品总数     */
	TotalItem *int64 `json:"total_item,omitempty" `

	/*
	   总页数     */
	TotalPage *int64 `json:"total_page,omitempty" `

	/*
	   埋点信息     */
	TrackParams *TaobaoJuItemsSearchTrackparams `json:"track_params,omitempty" `
}

func (s *TaobaoJuItemsSearchPaginationResult) SetCurrentPage(v int64) *TaobaoJuItemsSearchPaginationResult {
	s.CurrentPage = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetExtend(v TaobaoJuItemsSearchExtend) *TaobaoJuItemsSearchPaginationResult {
	s.Extend = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetModelList(v []TaobaoJuItemsSearchItems) *TaobaoJuItemsSearchPaginationResult {
	s.ModelList = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetMsgCode(v string) *TaobaoJuItemsSearchPaginationResult {
	s.MsgCode = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetMsgInfo(v string) *TaobaoJuItemsSearchPaginationResult {
	s.MsgInfo = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetPageSize(v int64) *TaobaoJuItemsSearchPaginationResult {
	s.PageSize = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetSuccess(v bool) *TaobaoJuItemsSearchPaginationResult {
	s.Success = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetTotalItem(v int64) *TaobaoJuItemsSearchPaginationResult {
	s.TotalItem = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetTotalPage(v int64) *TaobaoJuItemsSearchPaginationResult {
	s.TotalPage = &v
	return s
}
func (s *TaobaoJuItemsSearchPaginationResult) SetTrackParams(v TaobaoJuItemsSearchTrackparams) *TaobaoJuItemsSearchPaginationResult {
	s.TrackParams = &v
	return s
}
