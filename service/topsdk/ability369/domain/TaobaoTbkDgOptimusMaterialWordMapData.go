package domain

type TaobaoTbkDgOptimusMaterialWordMapData struct {
	/*
	   链接-商品相关关联词落地页地址     */
	Url *string `json:"url,omitempty" `

	/*
	   商品相关的关联词     */
	Word *string `json:"word,omitempty" `
}

func (s *TaobaoTbkDgOptimusMaterialWordMapData) SetUrl(v string) *TaobaoTbkDgOptimusMaterialWordMapData {
	s.Url = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialWordMapData) SetWord(v string) *TaobaoTbkDgOptimusMaterialWordMapData {
	s.Word = &v
	return s
}
