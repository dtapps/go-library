package domain

type TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto struct {
	/*
	   解冻金额     */
	UnfreezeAmount *string `json:"unfreeze_amount,omitempty" `

	/*
	   解冻红包个数     */
	UnfreezeNum *int64 `json:"unfreeze_num,omitempty" `

	/*
	   失效回退金额     */
	RefundAmount *string `json:"refund_amount,omitempty" `

	/*
	   失效回退红包个数     */
	RefundNum *int64 `json:"refund_num,omitempty" `

	/*
	   引导预估佣金金额     */
	PreCommissionAmount *string `json:"pre_commission_amount,omitempty" `

	/*
	   引导成交金额     */
	AlipayAmount *string `json:"alipay_amount,omitempty" `

	/*
	   红包核销金额     */
	UseAmount *string `json:"use_amount,omitempty" `

	/*
	   红包核销个数     */
	UseNum *int64 `json:"use_num,omitempty" `

	/*
	   红包领取金额     */
	WinAmount *string `json:"win_amount,omitempty" `

	/*
	   红包领取个数     */
	WinNum *int64 `json:"win_num,omitempty" `

	/*
	   退款红包金额     */
	FpRefundAmount *string `json:"fp_refund_amount,omitempty" `

	/*
	   退款红包个数     */
	FpRefundNum *int64 `json:"fp_refund_num,omitempty" `
}

func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetUnfreezeAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.UnfreezeAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetUnfreezeNum(v int64) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.UnfreezeNum = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetRefundAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.RefundAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetRefundNum(v int64) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.RefundNum = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetPreCommissionAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.PreCommissionAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetAlipayAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.AlipayAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetUseAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.UseAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetUseNum(v int64) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.UseNum = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetWinAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.WinAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetWinNum(v int64) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.WinNum = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetFpRefundAmount(v string) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.FpRefundAmount = &v
	return s
}
func (s *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto) SetFpRefundNum(v int64) *TaobaoTbkDgVegasTljInstanceReportTljInstanceReportDto {
	s.FpRefundNum = &v
	return s
}
