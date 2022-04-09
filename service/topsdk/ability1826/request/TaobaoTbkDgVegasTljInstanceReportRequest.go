package request

type TaobaoTbkDgVegasTljInstanceReportRequest struct {
	/*
	   实例ID     */
	RightsId *string `json:"rights_id" required:"true" `
}

func (s *TaobaoTbkDgVegasTljInstanceReportRequest) SetRightsId(v string) *TaobaoTbkDgVegasTljInstanceReportRequest {
	s.RightsId = &v
	return s
}

func (req *TaobaoTbkDgVegasTljInstanceReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RightsId != nil {
		paramMap["rights_id"] = *req.RightsId
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasTljInstanceReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
