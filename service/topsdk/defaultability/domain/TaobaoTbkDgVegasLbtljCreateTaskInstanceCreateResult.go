package domain

type TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult struct {
	/*
	   直领淘礼金id     */
	RightsId *string `json:"rights_id,omitempty" `

	/*
	   发放url     */
	SendUrl *string `json:"send_url,omitempty" `

	/*
	   裂变淘礼金id     */
	TaskRightsId *string `json:"task_rights_id,omitempty" `

	/*
	   裂变任务id     */
	TaskId *string `json:"task_id,omitempty" `
}

func (s *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult) SetRightsId(v string) *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult {
	s.RightsId = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult) SetSendUrl(v string) *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult {
	s.SendUrl = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult) SetTaskRightsId(v string) *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult {
	s.TaskRightsId = &v
	return s
}
func (s *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult) SetTaskId(v string) *TaobaoTbkDgVegasLbtljCreateTaskInstanceCreateResult {
	s.TaskId = &v
	return s
}
