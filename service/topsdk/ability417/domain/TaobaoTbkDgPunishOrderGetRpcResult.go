package domain

type TaobaoTbkDgPunishOrderGetRpcResult struct {
	/*
	   结果     */
	Data *TaobaoTbkDgPunishOrderGetPageResult `json:"data,omitempty" `

	/*
	   业务出错的描述     */
	BizErrorDesc *string `json:"biz_error_desc,omitempty" `

	/*
	   业务出错的状态码     */
	BizErrorCode *int64 `json:"biz_error_code,omitempty" `

	/*
	   执行结果     */
	ResultMsg *string `json:"result_msg,omitempty" `

	/*
	   执行结果状态码     */
	ResultCode *int64 `json:"result_code,omitempty" `
}

func (s *TaobaoTbkDgPunishOrderGetRpcResult) SetData(v TaobaoTbkDgPunishOrderGetPageResult) *TaobaoTbkDgPunishOrderGetRpcResult {
	s.Data = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetRpcResult) SetBizErrorDesc(v string) *TaobaoTbkDgPunishOrderGetRpcResult {
	s.BizErrorDesc = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetRpcResult) SetBizErrorCode(v int64) *TaobaoTbkDgPunishOrderGetRpcResult {
	s.BizErrorCode = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetRpcResult) SetResultMsg(v string) *TaobaoTbkDgPunishOrderGetRpcResult {
	s.ResultMsg = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetRpcResult) SetResultCode(v int64) *TaobaoTbkDgPunishOrderGetRpcResult {
	s.ResultCode = &v
	return s
}
