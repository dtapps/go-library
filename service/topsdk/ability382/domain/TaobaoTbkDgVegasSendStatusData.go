package domain

type TaobaoTbkDgVegasSendStatusData struct {
	/*
	   返回结果封装对象     */
	ResultList *[]TaobaoTbkDgVegasSendStatusMapData `json:"result_list,omitempty" `
}

func (s *TaobaoTbkDgVegasSendStatusData) SetResultList(v []TaobaoTbkDgVegasSendStatusMapData) *TaobaoTbkDgVegasSendStatusData {
	s.ResultList = &v
	return s
}
