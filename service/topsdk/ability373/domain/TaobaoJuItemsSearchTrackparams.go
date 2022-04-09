package domain

type TaobaoJuItemsSearchTrackparams struct {
	/*
	   empty     */
	Empty *bool `json:"empty,omitempty" `
}

func (s *TaobaoJuItemsSearchTrackparams) SetEmpty(v bool) *TaobaoJuItemsSearchTrackparams {
	s.Empty = &v
	return s
}
