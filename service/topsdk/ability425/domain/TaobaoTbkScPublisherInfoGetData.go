package domain

type TaobaoTbkScPublisherInfoGetData struct {
	/*
	   共享字段 - 总记录数     */
	TotalCount *int64 `json:"total_count,omitempty" `

	/*
	   共享字段 - 渠道或会员列表     */
	InviterList *[]TaobaoTbkScPublisherInfoGetMapData `json:"inviter_list,omitempty" `

	/*
	   渠道专属pidList     */
	RootPidChannelList *[]string `json:"root_pid_channel_list,omitempty" `
}

func (s *TaobaoTbkScPublisherInfoGetData) SetTotalCount(v int64) *TaobaoTbkScPublisherInfoGetData {
	s.TotalCount = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetData) SetInviterList(v []TaobaoTbkScPublisherInfoGetMapData) *TaobaoTbkScPublisherInfoGetData {
	s.InviterList = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetData) SetRootPidChannelList(v []string) *TaobaoTbkScPublisherInfoGetData {
	s.RootPidChannelList = &v
	return s
}
