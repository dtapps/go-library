package request

type TaobaoTbkDgVegasSendReportRequest struct {
	/*
	   统计日期     */
	BizDate *string `json:"biz_date" required:"true" `
	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" required:"false" `
	/*
	   红包活动id：1462     */
	ActivityId *int64 `json:"activity_id" required:"true" `
	/*
	   页码 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   每页大小 defalutValue��10    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoTbkDgVegasSendReportRequest) SetBizDate(v string) *TaobaoTbkDgVegasSendReportRequest {
	s.BizDate = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetRelationId(v int64) *TaobaoTbkDgVegasSendReportRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetActivityId(v int64) *TaobaoTbkDgVegasSendReportRequest {
	s.ActivityId = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetPageNo(v int64) *TaobaoTbkDgVegasSendReportRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgVegasSendReportRequest) SetPageSize(v int64) *TaobaoTbkDgVegasSendReportRequest {
	s.PageSize = &v
	return s
}

func (req *TaobaoTbkDgVegasSendReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BizDate != nil {
		paramMap["biz_date"] = *req.BizDate
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.ActivityId != nil {
		paramMap["activity_id"] = *req.ActivityId
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasSendReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
