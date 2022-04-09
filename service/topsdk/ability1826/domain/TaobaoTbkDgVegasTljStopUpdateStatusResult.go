package domain

type TaobaoTbkDgVegasTljStopUpdateStatusResult struct {
	/*
	   暂停成功     */
	UpdateSuccess *bool `json:"update_success,omitempty" `
}

func (s *TaobaoTbkDgVegasTljStopUpdateStatusResult) SetUpdateSuccess(v bool) *TaobaoTbkDgVegasTljStopUpdateStatusResult {
	s.UpdateSuccess = &v
	return s
}
