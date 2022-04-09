package domain

type TaobaoTbkCouponGetMapData struct {
	/*
	   优惠券门槛金额     */
	CouponStartFee *string `json:"coupon_start_fee,omitempty" `

	/*
	   优惠券剩余量     */
	CouponRemainCount *int64 `json:"coupon_remain_count,omitempty" `

	/*
	   优惠券总量     */
	CouponTotalCount *int64 `json:"coupon_total_count,omitempty" `

	/*
	   优惠券结束时间     */
	CouponEndTime *string `json:"coupon_end_time,omitempty" `

	/*
	   优惠券开始时间     */
	CouponStartTime *string `json:"coupon_start_time,omitempty" `

	/*
	   优惠券金额     */
	CouponAmount *string `json:"coupon_amount,omitempty" `

	/*
	   券类型，1 表示全网公开券，4 表示妈妈渠道券     */
	CouponSrcScene *int64 `json:"coupon_src_scene,omitempty" `

	/*
	   券属性，0表示店铺券，1表示单品券     */
	CouponType *int64 `json:"coupon_type,omitempty" `

	/*
	   券ID     */
	CouponActivityId *string `json:"coupon_activity_id,omitempty" `
}

func (s *TaobaoTbkCouponGetMapData) SetCouponStartFee(v string) *TaobaoTbkCouponGetMapData {
	s.CouponStartFee = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponRemainCount(v int64) *TaobaoTbkCouponGetMapData {
	s.CouponRemainCount = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponTotalCount(v int64) *TaobaoTbkCouponGetMapData {
	s.CouponTotalCount = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponEndTime(v string) *TaobaoTbkCouponGetMapData {
	s.CouponEndTime = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponStartTime(v string) *TaobaoTbkCouponGetMapData {
	s.CouponStartTime = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponAmount(v string) *TaobaoTbkCouponGetMapData {
	s.CouponAmount = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponSrcScene(v int64) *TaobaoTbkCouponGetMapData {
	s.CouponSrcScene = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponType(v int64) *TaobaoTbkCouponGetMapData {
	s.CouponType = &v
	return s
}
func (s *TaobaoTbkCouponGetMapData) SetCouponActivityId(v string) *TaobaoTbkCouponGetMapData {
	s.CouponActivityId = &v
	return s
}
