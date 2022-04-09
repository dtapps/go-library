package domain

type TaobaoTbkDgOptimusMaterialFavoritesInfo struct {
	/*
	   选品库总数量     */
	TotalCount *int64 `json:"total_count,omitempty" `

	/*
	   选品库详细信息     */
	FavoritesList *[]TaobaoTbkDgOptimusMaterialFavoritesDetail `json:"favorites_list,omitempty" `
}

func (s *TaobaoTbkDgOptimusMaterialFavoritesInfo) SetTotalCount(v int64) *TaobaoTbkDgOptimusMaterialFavoritesInfo {
	s.TotalCount = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialFavoritesInfo) SetFavoritesList(v []TaobaoTbkDgOptimusMaterialFavoritesDetail) *TaobaoTbkDgOptimusMaterialFavoritesInfo {
	s.FavoritesList = &v
	return s
}
