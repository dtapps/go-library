package request

import (
	"topsdk/util"
)

type TaobaoTbkDgNewuserOrderGetRequest struct {
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   页码，默认1 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   mm_xxx_xxx_xxx的第三位     */
	AdzoneId *int64 `json:"adzone_id,omitempty" required:"false" `
	/*
	   开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" required:"false" `
	/*
	   结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" required:"false" `
	/*
	   活动id， 活动名称与活动ID列表（该字段已废弃）     */
	ActivityId *string `json:"activity_id" required:"true" `
}

func (s *TaobaoTbkDgNewuserOrderGetRequest) SetPageSize(v int64) *TaobaoTbkDgNewuserOrderGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetRequest) SetPageNo(v int64) *TaobaoTbkDgNewuserOrderGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetRequest) SetAdzoneId(v int64) *TaobaoTbkDgNewuserOrderGetRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetRequest) SetStartTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetRequest {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetRequest) SetEndTime(v util.LocalTime) *TaobaoTbkDgNewuserOrderGetRequest {
	s.EndTime = &v
	return s
}
func (s *TaobaoTbkDgNewuserOrderGetRequest) SetActivityId(v string) *TaobaoTbkDgNewuserOrderGetRequest {
	s.ActivityId = &v
	return s
}

func (req *TaobaoTbkDgNewuserOrderGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.StartTime != nil {
		paramMap["start_time"] = *req.StartTime
	}
	if req.EndTime != nil {
		paramMap["end_time"] = *req.EndTime
	}
	if req.ActivityId != nil {
		paramMap["activity_id"] = *req.ActivityId
	}
	return paramMap
}

func (req *TaobaoTbkDgNewuserOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
