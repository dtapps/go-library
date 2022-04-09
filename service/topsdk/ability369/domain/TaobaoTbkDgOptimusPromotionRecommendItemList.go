package domain

type TaobaoTbkDgOptimusPromotionRecommendItemList struct {
	/*
	   权益推荐商品id     */
	ItemId *int64 `json:"item_id,omitempty" `

	/*
	   商品链接     */
	Url *string `json:"url,omitempty" `
}

func (s *TaobaoTbkDgOptimusPromotionRecommendItemList) SetItemId(v int64) *TaobaoTbkDgOptimusPromotionRecommendItemList {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionRecommendItemList) SetUrl(v string) *TaobaoTbkDgOptimusPromotionRecommendItemList {
	s.Url = &v
	return s
}
