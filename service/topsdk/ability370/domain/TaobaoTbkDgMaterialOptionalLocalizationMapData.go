package domain

type TaobaoTbkDgMaterialOptionalLocalizationMapData struct {
	/*
	   本地化-配送时间     */
	OrderLeadTime *string `json:"order_lead_time,omitempty" `

	/*
	   本地化-用户评分     */
	UserRating *string `json:"user_rating,omitempty" `

	/*
	   本地化-起送价     */
	DeliveryMinPrice *string `json:"delivery_min_price,omitempty" `

	/*
	   本地化-配送费     */
	DeliveryFee *string `json:"delivery_fee,omitempty" `

	/*
	   本地化-配送费原价     */
	OriginalDeliveryFee *string `json:"original_delivery_fee,omitempty" `

	/*
	   本地化-配送类型；0：蜂鸟专送 1：蜂鸟快送 2：商家自配 3: 全城送     */
	DeliveryType *string `json:"delivery_type,omitempty" `

	/*
	   本地化-推荐理由     */
	RecommendReasons *[]string `json:"recommend_reasons,omitempty" `

	/*
	   本地化-营销标签     */
	SaleTags *[]string `json:"sale_tags,omitempty" `

	/*
	   本地化-单店商品列表     */
	FoodItemList *[]TaobaoTbkDgMaterialOptionalFoodMapData `json:"food_item_list,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetOrderLeadTime(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.OrderLeadTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetUserRating(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.UserRating = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetDeliveryMinPrice(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.DeliveryMinPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetDeliveryFee(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.DeliveryFee = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetOriginalDeliveryFee(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.OriginalDeliveryFee = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetDeliveryType(v string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.DeliveryType = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetRecommendReasons(v []string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.RecommendReasons = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetSaleTags(v []string) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.SaleTags = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalLocalizationMapData) SetFoodItemList(v []TaobaoTbkDgMaterialOptionalFoodMapData) *TaobaoTbkDgMaterialOptionalLocalizationMapData {
	s.FoodItemList = &v
	return s
}
