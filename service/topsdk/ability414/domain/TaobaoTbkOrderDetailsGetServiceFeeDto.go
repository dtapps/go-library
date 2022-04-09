package domain

type TaobaoTbkOrderDetailsGetServiceFeeDto struct {
	/*
	   专项服务费来源，122-渠道     */
	TkShareRoleType *int64 `json:"tk_share_role_type,omitempty" `

	/*
	   专项服务费率     */
	ShareRelativeRate *string `json:"share_relative_rate,omitempty" `

	/*
	   结算专项服务费     */
	ShareFee *string `json:"share_fee,omitempty" `

	/*
	   预估专项服务费     */
	SharePreFee *string `json:"share_pre_fee,omitempty" `
}

func (s *TaobaoTbkOrderDetailsGetServiceFeeDto) SetTkShareRoleType(v int64) *TaobaoTbkOrderDetailsGetServiceFeeDto {
	s.TkShareRoleType = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDto) SetShareRelativeRate(v string) *TaobaoTbkOrderDetailsGetServiceFeeDto {
	s.ShareRelativeRate = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDto) SetShareFee(v string) *TaobaoTbkOrderDetailsGetServiceFeeDto {
	s.ShareFee = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDto) SetSharePreFee(v string) *TaobaoTbkOrderDetailsGetServiceFeeDto {
	s.SharePreFee = &v
	return s
}
