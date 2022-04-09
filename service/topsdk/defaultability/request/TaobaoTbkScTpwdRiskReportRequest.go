package request

type TaobaoTbkScTpwdRiskReportRequest struct {
	/*
	   如有pid，则查询pid下的relationid名单；如没有pid，则查询授权userid下的pid名单     */
	Pid *string `json:"pid,omitempty" required:"false" `
	/*
	   分页参数     */
	Offset *int32 `json:"offset" required:"true" `
	/*
	   分页参数，一次最多不能超过1000     */
	Limit *int32 `json:"limit" required:"true" `
}

func (s *TaobaoTbkScTpwdRiskReportRequest) SetPid(v string) *TaobaoTbkScTpwdRiskReportRequest {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkScTpwdRiskReportRequest) SetOffset(v int32) *TaobaoTbkScTpwdRiskReportRequest {
	s.Offset = &v
	return s
}
func (s *TaobaoTbkScTpwdRiskReportRequest) SetLimit(v int32) *TaobaoTbkScTpwdRiskReportRequest {
	s.Limit = &v
	return s
}

func (req *TaobaoTbkScTpwdRiskReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.Offset != nil {
		paramMap["offset"] = *req.Offset
	}
	if req.Limit != nil {
		paramMap["limit"] = *req.Limit
	}
	return paramMap
}

func (req *TaobaoTbkScTpwdRiskReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
