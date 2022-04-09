package request

type TaobaoTbkDgCpaActivityReportRequest struct {
	/*
	   CPA活动ID，详见https://www.yuque.com/docs/share/16905f3f-3a22-4e7c-b8c3-4d23791d42f7?#     */
	EventId *int64 `json:"event_id" required:"true" `
	/*
	   日期(yyyyMMdd)     */
	BizDate *string `json:"biz_date" required:"true" `
	/*
	   分页页数，从1开始     */
	PageNo *int32 `json:"page_no,omitempty" required:"false" `
	/*
	   数据类型：数据类型:1预估 2结算 （选择1可查询含当天实时预估统计的累计数据，选择2可查询最晚截止昨天结算的累计数据，具体逻辑以活动规则描述为准；） defalutValue��1    */
	QueryType *int32 `json:"query_type,omitempty" required:"false" `
	/*
	   分页大小     */
	PageSize *int32 `json:"page_size,omitempty" required:"false" `
	/*
	   媒体三段式id（如果传入pid则返回pid汇总数据，不传则返回member维度统计数据，pid和relationid不可同时传入）     */
	Pid *string `json:"pid,omitempty" required:"false" `
	/*
	   代理id（如果传入rid则返回rid统计数据，不传则返回member维度统计数据）     */
	RelationId *int64 `json:"relation_id,omitempty" required:"false" `
}

func (s *TaobaoTbkDgCpaActivityReportRequest) SetEventId(v int64) *TaobaoTbkDgCpaActivityReportRequest {
	s.EventId = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetBizDate(v string) *TaobaoTbkDgCpaActivityReportRequest {
	s.BizDate = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetPageNo(v int32) *TaobaoTbkDgCpaActivityReportRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetQueryType(v int32) *TaobaoTbkDgCpaActivityReportRequest {
	s.QueryType = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetPageSize(v int32) *TaobaoTbkDgCpaActivityReportRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetPid(v string) *TaobaoTbkDgCpaActivityReportRequest {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportRequest) SetRelationId(v int64) *TaobaoTbkDgCpaActivityReportRequest {
	s.RelationId = &v
	return s
}

func (req *TaobaoTbkDgCpaActivityReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.EventId != nil {
		paramMap["event_id"] = *req.EventId
	}
	if req.BizDate != nil {
		paramMap["biz_date"] = *req.BizDate
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.QueryType != nil {
		paramMap["query_type"] = *req.QueryType
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	return paramMap
}

func (req *TaobaoTbkDgCpaActivityReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
