package domain

type TaobaoTbkDgCpaActivityReportVegasCpaReportDTO struct {
	/*
	   符合奖励要求的累计用户数；按入参是预估/结算，区分用户数为预估or可结算结果；     */
	Union30dLxUv *int64 `json:"union_30d_lx_uv,omitempty" `

	/*
	   奖励金额；按入参是预估/结算，区分获得金额为预估or可结算结果；     */
	RewardAmount *string `json:"reward_amount,omitempty" `

	/*
	   rid，当查询数据为rid维度时返回该字段     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   统计日期（统计活动期内，截止 统计日期 的数据）     */
	BizDate *string `json:"biz_date,omitempty" `

	/*
	   媒体三段式id，当查询数据为pid维度时返回该字段     */
	Pid *string `json:"pid,omitempty" `

	/*
	   数据类型:1预估 2结算     */
	QueryType *int32 `json:"query_type,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetUnion30dLxUv(v int64) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.Union30dLxUv = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetRewardAmount(v string) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.RewardAmount = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetRelationId(v int64) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetBizDate(v string) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.BizDate = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetPid(v string) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO) SetQueryType(v int32) *TaobaoTbkDgCpaActivityReportVegasCpaReportDTO {
	s.QueryType = &v
	return s
}
