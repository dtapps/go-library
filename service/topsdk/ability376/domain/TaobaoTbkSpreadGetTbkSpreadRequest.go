package domain

type TaobaoTbkSpreadGetTbkSpreadRequest struct {
	/*
	   原始url, 只支持uland.taobao.com，s.click.taobao.com， ai.taobao.com，temai.taobao.com的域名转换，否则判错     */
	Url *string `json:"url,omitempty" `
}

func (s *TaobaoTbkSpreadGetTbkSpreadRequest) SetUrl(v string) *TaobaoTbkSpreadGetTbkSpreadRequest {
	s.Url = &v
	return s
}
