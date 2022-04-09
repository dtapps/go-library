package request

type TaobaoTbkTpwdCreateRequest struct {
	/*
	   兼容旧版本api参数，无实际作用     */
	Text *string `json:"text,omitempty" required:"false" `
	/*
	   兼容旧版本api参数，无实际作用     */
	Logo *string `json:"logo,omitempty" required:"false" `
	/*
	   兼容旧版本api参数，无实际作用     */
	Ext *string `json:"ext,omitempty" required:"false" `
	/*
	   兼容旧版本api参数，无实际作用     */
	UserId *string `json:"user_id,omitempty" required:"false" `
	/*
	   联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令     */
	Url *string `json:"url" required:"true" `
}

func (s *TaobaoTbkTpwdCreateRequest) SetText(v string) *TaobaoTbkTpwdCreateRequest {
	s.Text = &v
	return s
}
func (s *TaobaoTbkTpwdCreateRequest) SetLogo(v string) *TaobaoTbkTpwdCreateRequest {
	s.Logo = &v
	return s
}
func (s *TaobaoTbkTpwdCreateRequest) SetExt(v string) *TaobaoTbkTpwdCreateRequest {
	s.Ext = &v
	return s
}
func (s *TaobaoTbkTpwdCreateRequest) SetUserId(v string) *TaobaoTbkTpwdCreateRequest {
	s.UserId = &v
	return s
}
func (s *TaobaoTbkTpwdCreateRequest) SetUrl(v string) *TaobaoTbkTpwdCreateRequest {
	s.Url = &v
	return s
}

func (req *TaobaoTbkTpwdCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Text != nil {
		paramMap["text"] = *req.Text
	}
	if req.Logo != nil {
		paramMap["logo"] = *req.Logo
	}
	if req.Ext != nil {
		paramMap["ext"] = *req.Ext
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.Url != nil {
		paramMap["url"] = *req.Url
	}
	return paramMap
}

func (req *TaobaoTbkTpwdCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
