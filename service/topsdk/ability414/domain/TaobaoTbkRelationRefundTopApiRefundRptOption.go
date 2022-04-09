package domain

import (
	"topsdk/util"
)

type TaobaoTbkRelationRefundTopApiRefundRptOption struct {
	/*
	   pagesize     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间     */
	SearchType *int64 `json:"search_type,omitempty" `

	/*
	   1 表示2方，2表示3方，0表示不限     */
	RefundType *int64 `json:"refund_type,omitempty" `

	/*
	   开始时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   pagenumber     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   1代表渠道关系id，2代表会员关系id     */
	BizType *int64 `json:"biz_type,omitempty" `
}

func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetPageSize(v int64) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetSearchType(v int64) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.SearchType = &v
	return s
}
func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetRefundType(v int64) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.RefundType = &v
	return s
}
func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetStartTime(v util.LocalTime) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetPageNo(v int64) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkRelationRefundTopApiRefundRptOption) SetBizType(v int64) *TaobaoTbkRelationRefundTopApiRefundRptOption {
	s.BizType = &v
	return s
}
