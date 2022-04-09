package domain

type TaobaoTbkDgNewuserOrderGetOrderData struct {
	/*
	   预估佣金     */
	Commission *string `json:"commission,omitempty" `

	/*
	   收货时间     */
	ConfirmReceiveTime *string `json:"confirm_receive_time,omitempty" `

	/*
	   支付时间     */
	PayTime *string `json:"pay_time,omitempty" `

	/*
	   订单号     */
	OrderNo *string `json:"order_no,omitempty" `
}

func (s *TaobaoTbkDgNewuserOrderGetOrderData) SetCommission(v string) *TaobaoTbkDgNewuserOrderGetOrderData {
	s.Commission = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetOrderData) SetConfirmReceiveTime(v string) *TaobaoTbkDgNewuserOrderGetOrderData {
	s.ConfirmReceiveTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetOrderData) SetPayTime(v string) *TaobaoTbkDgNewuserOrderGetOrderData {
	s.PayTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetOrderData) SetOrderNo(v string) *TaobaoTbkDgNewuserOrderGetOrderData {
	s.OrderNo = &v
	return s
}
