package domain

type TaobaoTbkDgOptimusMaterialFavoritesDetail struct {
	/*
	   选品库id     */
	FavoritesId *int64 `json:"favorites_id,omitempty" `

	/*
	   选品库标题     */
	FavoritesTitle *string `json:"favorites_title,omitempty" `
}

func (s *TaobaoTbkDgOptimusMaterialFavoritesDetail) SetFavoritesId(v int64) *TaobaoTbkDgOptimusMaterialFavoritesDetail {
	s.FavoritesId = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialFavoritesDetail) SetFavoritesTitle(v string) *TaobaoTbkDgOptimusMaterialFavoritesDetail {
	s.FavoritesTitle = &v
	return s
}
