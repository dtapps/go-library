package domain

type TaobaoTbkDgMaterialOptionalMaifanPromotionDTO struct {
	/*
	   猫超买返卡活动结束时间     */
	MaifanPromotionEndTime *string `json:"maifan_promotion_end_time,omitempty" `

	/*
	   猫超买返卡活动开始时间     */
	MaifanPromotionStartTime *string `json:"maifan_promotion_start_time,omitempty" `

	/*
	   猫超买返卡面额     */
	MaifanPromotionDiscount *string `json:"maifan_promotion_discount,omitempty" `

	/*
	   猫超买返卡总数，-1代表不限量，其他大于等于0的值为总数     */
	MaifanPromotionCondition *string `json:"maifan_promotion_condition,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionEndTime(v string) *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO {
	s.MaifanPromotionEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionStartTime(v string) *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO {
	s.MaifanPromotionStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionDiscount(v string) *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO {
	s.MaifanPromotionDiscount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionCondition(v string) *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO {
	s.MaifanPromotionCondition = &v
	return s
}
