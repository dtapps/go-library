package request

type TaobaoTbkDgVegasTljStopRequest struct {
	/*
	   adzoneId     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   创建淘礼金时返回的rightsId     */
	RightsId *string `json:"rights_id" required:"true" `
}

func (s *TaobaoTbkDgVegasTljStopRequest) SetAdzoneId(v int64) *TaobaoTbkDgVegasTljStopRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgVegasTljStopRequest) SetRightsId(v string) *TaobaoTbkDgVegasTljStopRequest {
	s.RightsId = &v
	return s
}

func (req *TaobaoTbkDgVegasTljStopRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.RightsId != nil {
		paramMap["rights_id"] = *req.RightsId
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasTljStopRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
