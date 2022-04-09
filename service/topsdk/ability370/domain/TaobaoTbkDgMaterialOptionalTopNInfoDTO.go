package domain

type TaobaoTbkDgMaterialOptionalTopNInfoDTO struct {
	/*
	   前N件剩余库存     */
	TopnQuantity *int64 `json:"topn_quantity,omitempty" `

	/*
	   前N件初始总库存     */
	TopnTotalCount *int64 `json:"topn_total_count,omitempty" `

	/*
	   前N件佣金结束时间     */
	TopnEndTime *string `json:"topn_end_time,omitempty" `

	/*
	   前N件佣金开始时间     */
	TopnStartTime *string `json:"topn_start_time,omitempty" `

	/*
	   前N件佣金率     */
	TopnRate *string `json:"topn_rate,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalTopNInfoDTO) SetTopnQuantity(v int64) *TaobaoTbkDgMaterialOptionalTopNInfoDTO {
	s.TopnQuantity = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalTopNInfoDTO) SetTopnTotalCount(v int64) *TaobaoTbkDgMaterialOptionalTopNInfoDTO {
	s.TopnTotalCount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalTopNInfoDTO) SetTopnEndTime(v string) *TaobaoTbkDgMaterialOptionalTopNInfoDTO {
	s.TopnEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalTopNInfoDTO) SetTopnStartTime(v string) *TaobaoTbkDgMaterialOptionalTopNInfoDTO {
	s.TopnStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalTopNInfoDTO) SetTopnRate(v string) *TaobaoTbkDgMaterialOptionalTopNInfoDTO {
	s.TopnRate = &v
	return s
}
