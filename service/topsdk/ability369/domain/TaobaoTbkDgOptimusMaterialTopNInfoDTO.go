package domain

type TaobaoTbkDgOptimusMaterialTopNInfoDTO struct {
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

func (s *TaobaoTbkDgOptimusMaterialTopNInfoDTO) SetTopnQuantity(v int64) *TaobaoTbkDgOptimusMaterialTopNInfoDTO {
	s.TopnQuantity = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialTopNInfoDTO) SetTopnTotalCount(v int64) *TaobaoTbkDgOptimusMaterialTopNInfoDTO {
	s.TopnTotalCount = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialTopNInfoDTO) SetTopnEndTime(v string) *TaobaoTbkDgOptimusMaterialTopNInfoDTO {
	s.TopnEndTime = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialTopNInfoDTO) SetTopnStartTime(v string) *TaobaoTbkDgOptimusMaterialTopNInfoDTO {
	s.TopnStartTime = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialTopNInfoDTO) SetTopnRate(v string) *TaobaoTbkDgOptimusMaterialTopNInfoDTO {
	s.TopnRate = &v
	return s
}
