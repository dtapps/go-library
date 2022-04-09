package domain

type TaobaoTbkDgNewuserOrderGetResults struct {
	/*
	   data     */
	Data *TaobaoTbkDgNewuserOrderGetData `json:"data,omitempty" `
}

func (s *TaobaoTbkDgNewuserOrderGetResults) SetData(v TaobaoTbkDgNewuserOrderGetData) *TaobaoTbkDgNewuserOrderGetResults {
	s.Data = &v
	return s
}
