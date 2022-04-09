package domain

type TaobaoJuItemsSearchExtend struct {
	/*
	   empty     */
	Empty *bool `json:"empty,omitempty" `
}

func (s *TaobaoJuItemsSearchExtend) SetEmpty(v bool) *TaobaoJuItemsSearchExtend {
	s.Empty = &v
	return s
}
