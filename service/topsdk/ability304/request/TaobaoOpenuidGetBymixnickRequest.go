package request

type TaobaoOpenuidGetBymixnickRequest struct {
	/*
	   无线类应用获取到的混淆的nick     */
	MixNick *string `json:"mix_nick" required:"true" `
}

func (s *TaobaoOpenuidGetBymixnickRequest) SetMixNick(v string) *TaobaoOpenuidGetBymixnickRequest {
	s.MixNick = &v
	return s
}

func (req *TaobaoOpenuidGetBymixnickRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.MixNick != nil {
		paramMap["mix_nick"] = *req.MixNick
	}
	return paramMap
}

func (req *TaobaoOpenuidGetBymixnickRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
