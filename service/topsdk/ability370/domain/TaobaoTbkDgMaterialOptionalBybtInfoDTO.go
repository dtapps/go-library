package domain

type TaobaoTbkDgMaterialOptionalBybtInfoDTO struct {
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

func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtBrandLogo(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtBrandLogo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtPicUrl(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtPicUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtItemTags(v []string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtItemTags = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtCouponAmount(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtCouponAmount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtShowPrice(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtShowPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtLowestPrice(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtLowestPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtEndTime(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalBybtInfoDTO) SetBybtStartTime(v string) *TaobaoTbkDgMaterialOptionalBybtInfoDTO {
	s.BybtStartTime = &v
	return s
}
