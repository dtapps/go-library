package domain

type TaobaoTbkDgOptimusPromotionPromotionExtend struct {
	/*
	   权益推荐商品     */
	RecommendItemList *[]TaobaoTbkDgOptimusPromotionRecommendItemList `json:"recommend_item_list,omitempty" `

	/*
	   有价券信息     */
	YoujiaCouponInfo *TaobaoTbkDgOptimusPromotionYoujiacouponinfo `json:"youjia_coupon_info,omitempty" `

	/*
	   权益链接     */
	PromotionUrl *string `json:"promotion_url,omitempty" `
}

func (s *TaobaoTbkDgOptimusPromotionPromotionExtend) SetRecommendItemList(v []TaobaoTbkDgOptimusPromotionRecommendItemList) *TaobaoTbkDgOptimusPromotionPromotionExtend {
	s.RecommendItemList = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionPromotionExtend) SetYoujiaCouponInfo(v TaobaoTbkDgOptimusPromotionYoujiacouponinfo) *TaobaoTbkDgOptimusPromotionPromotionExtend {
	s.YoujiaCouponInfo = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionPromotionExtend) SetPromotionUrl(v string) *TaobaoTbkDgOptimusPromotionPromotionExtend {
	s.PromotionUrl = &v
	return s
}
