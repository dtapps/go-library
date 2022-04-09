package domain

type TaobaoTbkDgOptimusMaterialBybtInfoDTO struct {
	/*
	   百亿补贴品牌logo     */
	BybtBrandLogo *string `json:"bybt_brand_logo,omitempty" `

	/*
	   百亿补贴白底图     */
	BybtPicUrl *string `json:"bybt_pic_url,omitempty" `

	/*
	   百亿补贴商品特征标签，eg.今日发货、晚发补偿、限购一件等     */
	BybtItemTags *[]string `json:"bybt_item_tags,omitempty" `

	/*
	   百亿补贴专属券面额，仅限百亿补贴场景透出     */
	BybtCouponAmount *string `json:"bybt_coupon_amount,omitempty" `

	/*
	   百亿补贴页面实时价     */
	BybtShowPrice *string `json:"bybt_show_price,omitempty" `

	/*
	   全网对比参考价格     */
	BybtLowestPrice *string `json:"bybt_lowest_price,omitempty" `

	/*
	   商品的百亿补贴开始时间     */
	BybtEndTime *string `json:"bybt_end_time,omitempty" `

	/*
	   商品的百亿补贴结束时间     */
	BybtStartTime *string `json:"bybt_start_time,omitempty" `
}

func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtBrandLogo(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtBrandLogo = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtPicUrl(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtPicUrl = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtItemTags(v []string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtItemTags = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtCouponAmount(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtCouponAmount = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtShowPrice(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtShowPrice = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtLowestPrice(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtLowestPrice = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtEndTime(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtEndTime = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialBybtInfoDTO) SetBybtStartTime(v string) *TaobaoTbkDgOptimusMaterialBybtInfoDTO {
	s.BybtStartTime = &v
	return s
}
