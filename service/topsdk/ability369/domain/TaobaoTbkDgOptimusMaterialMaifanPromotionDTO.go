package domain

type TaobaoTbkDgOptimusMaterialMaifanPromotionDTO struct {
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

func (s *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO) SetMaifanPromotionEndTime(v string) *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO {
	s.MaifanPromotionEndTime = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO) SetMaifanPromotionStartTime(v string) *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO {
	s.MaifanPromotionStartTime = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO) SetMaifanPromotionDiscount(v string) *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO {
	s.MaifanPromotionDiscount = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO) SetMaifanPromotionCondition(v string) *TaobaoTbkDgOptimusMaterialMaifanPromotionDTO {
	s.MaifanPromotionCondition = &v
	return s
}
