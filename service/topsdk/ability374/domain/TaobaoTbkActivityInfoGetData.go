package domain

type TaobaoTbkActivityInfoGetData struct {
	/*
	   【本地化】微信推广二维码地址     */
	WxQrcodeUrl *string `json:"wx_qrcode_url,omitempty" `

	/*
	   淘客推广长链     */
	ClickUrl *string `json:"click_url,omitempty" `

	/*
	   淘客推广短链     */
	ShortClickUrl *string `json:"short_click_url,omitempty" `

	/*
	   投放平台, 1-PC 2-无线     */
	TerminalType *string `json:"terminal_type,omitempty" `

	/*
	   物料素材下载地址     */
	MaterialOssUrl *string `json:"material_oss_url,omitempty" `

	/*
	   会场名称     */
	PageName *string `json:"page_name,omitempty" `

	/*
	   活动开始时间     */
	PageStartTime *string `json:"page_start_time,omitempty" `

	/*
	   活动结束时间     */
	PageEndTime *string `json:"page_end_time,omitempty" `

	/*
	   【本地化】微信小程序路径     */
	WxMiniprogramPath *string `json:"wx_miniprogram_path,omitempty" `
}

func (s *TaobaoTbkActivityInfoGetData) SetWxQrcodeUrl(v string) *TaobaoTbkActivityInfoGetData {
	s.WxQrcodeUrl = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetClickUrl(v string) *TaobaoTbkActivityInfoGetData {
	s.ClickUrl = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetShortClickUrl(v string) *TaobaoTbkActivityInfoGetData {
	s.ShortClickUrl = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetTerminalType(v string) *TaobaoTbkActivityInfoGetData {
	s.TerminalType = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetMaterialOssUrl(v string) *TaobaoTbkActivityInfoGetData {
	s.MaterialOssUrl = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetPageName(v string) *TaobaoTbkActivityInfoGetData {
	s.PageName = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetPageStartTime(v string) *TaobaoTbkActivityInfoGetData {
	s.PageStartTime = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetPageEndTime(v string) *TaobaoTbkActivityInfoGetData {
	s.PageEndTime = &v
	return s
}
func (s *TaobaoTbkActivityInfoGetData) SetWxMiniprogramPath(v string) *TaobaoTbkActivityInfoGetData {
	s.WxMiniprogramPath = &v
	return s
}
