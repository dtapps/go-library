package domain

type TaobaoTbkScInvitecodeGetData struct {
	/*
	   邀请码     */
	InviterCode *string `json:"inviter_code,omitempty" `
}

func (s *TaobaoTbkScInvitecodeGetData) SetInviterCode(v string) *TaobaoTbkScInvitecodeGetData {
	s.InviterCode = &v
	return s
}
