package request

import (
	"topsdk/util"
)

type TaobaoFilesGetRequest struct {
	/*
	   下载链接状态。1:未下载。2:已下载     */
	Status *int64 `json:"status,omitempty" required:"false" `
	/*
	   搜索开始时间     */
	StartDate *util.LocalTime `json:"start_date" required:"true" `
	/*
	   搜索结束时间     */
	EndDate *util.LocalTime `json:"end_date" required:"true" `
}

func (s *TaobaoFilesGetRequest) SetStatus(v int64) *TaobaoFilesGetRequest {
	s.Status = &v
	return s
}
func (s *TaobaoFilesGetRequest) SetStartDate(v util.LocalTime) *TaobaoFilesGetRequest {
	s.StartDate = &v
	return s
}
func (s *TaobaoFilesGetRequest) SetEndDate(v util.LocalTime) *TaobaoFilesGetRequest {
	s.EndDate = &v
	return s
}

func (req *TaobaoFilesGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.StartDate != nil {
		paramMap["start_date"] = *req.StartDate
	}
	if req.EndDate != nil {
		paramMap["end_date"] = *req.EndDate
	}
	return paramMap
}

func (req *TaobaoFilesGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
