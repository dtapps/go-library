package domain

type TaobaoTbkDgMaterialOptionalFoodMapData struct {
	/*
	   本地化-商品图片     */
	FoodPic *string `json:"food_pic,omitempty" `

	/*
	   本地化-商品标题     */
	FoodTitle *string `json:"food_title,omitempty" `

	/*
	   本地化-商品促销价     */
	FoodPromotionPrice *string `json:"food_promotion_price,omitempty" `

	/*
	   本地化-商品原价     */
	FoodReservePrice *string `json:"food_reserve_price,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalFoodMapData) SetFoodPic(v string) *TaobaoTbkDgMaterialOptionalFoodMapData {
	s.FoodPic = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalFoodMapData) SetFoodTitle(v string) *TaobaoTbkDgMaterialOptionalFoodMapData {
	s.FoodTitle = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalFoodMapData) SetFoodPromotionPrice(v string) *TaobaoTbkDgMaterialOptionalFoodMapData {
	s.FoodPromotionPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalFoodMapData) SetFoodReservePrice(v string) *TaobaoTbkDgMaterialOptionalFoodMapData {
	s.FoodReservePrice = &v
	return s
}
