package request

type TaobaoTbkDgVegasTljReportRequest struct {
	/*
	   adzoneId     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   创建淘礼金时返回的rightsId     */
	RightsId *string `json:"rights_id" required:"true" `
}

func (s *TaobaoTbkDgVegasTljReportRequest) SetAdzoneId(v int64) *TaobaoTbkDgVegasTljReportRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgVegasTljReportRequest) SetRightsId(v string) *TaobaoTbkDgVegasTljReportRequest {
	s.RightsId = &v
	return s
}

func (req *TaobaoTbkDgVegasTljReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.RightsId != nil {
		paramMap["rights_id"] = *req.RightsId
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasTljReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
