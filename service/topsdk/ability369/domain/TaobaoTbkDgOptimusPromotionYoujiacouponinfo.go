package domain

type TaobaoTbkDgOptimusPromotionYoujiacouponinfo struct {
	/*
	   有价券商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品链接     */
	Url *string `json:"url,omitempty" `
}

func (s *TaobaoTbkDgOptimusPromotionYoujiacouponinfo) SetItemId(v string) *TaobaoTbkDgOptimusPromotionYoujiacouponinfo {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionYoujiacouponinfo) SetUrl(v string) *TaobaoTbkDgOptimusPromotionYoujiacouponinfo {
	s.Url = &v
	return s
}
