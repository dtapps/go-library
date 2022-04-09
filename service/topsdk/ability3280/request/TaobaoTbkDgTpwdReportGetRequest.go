package request

type TaobaoTbkDgTpwdReportGetRequest struct {
	/*
	   待查询的口令     */
	TaoPassword *string `json:"tao_password" required:"true" `
	/*
	   mm_xxx_xxx_xxx的第3段数字     */
	AdzoneId *string `json:"adzone_id" required:"true" `
}

func (s *TaobaoTbkDgTpwdReportGetRequest) SetTaoPassword(v string) *TaobaoTbkDgTpwdReportGetRequest {
	s.TaoPassword = &v
	return s
}
func (s *TaobaoTbkDgTpwdReportGetRequest) SetAdzoneId(v string) *TaobaoTbkDgTpwdReportGetRequest {
	s.AdzoneId = &v
	return s
}

func (req *TaobaoTbkDgTpwdReportGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.TaoPassword != nil {
		paramMap["tao_password"] = *req.TaoPassword
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	return paramMap
}

func (req *TaobaoTbkDgTpwdReportGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
