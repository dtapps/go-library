package domain

type TaobaoTbkRelationRefundRpcResult struct {
	/*
	   返回信息     */
	ResultMsg *string `json:"result_msg,omitempty" `

	/*
	   真正的业务数据结构     */
	Data *TaobaoTbkRelationRefundPageResult `json:"data,omitempty" `

	/*
	   接口返回值信息，跟rpc架构保持一致     */
	ResultCode *int64 `json:"result_code,omitempty" `

	/*
	   业务错误信息     */
	BizErrorDesc *string `json:"biz_error_desc,omitempty" `

	/*
	   业务错误码 101, 102,103     */
	BizErrorCode *int64 `json:"biz_error_code,omitempty" `
}

func (s *TaobaoTbkRelationRefundRpcResult) SetResultMsg(v string) *TaobaoTbkRelationRefundRpcResult {
	s.ResultMsg = &v
	return s
}
func (s *TaobaoTbkRelationRefundRpcResult) SetData(v TaobaoTbkRelationRefundPageResult) *TaobaoTbkRelationRefundRpcResult {
	s.Data = &v
	return s
}
func (s *TaobaoTbkRelationRefundRpcResult) SetResultCode(v int64) *TaobaoTbkRelationRefundRpcResult {
	s.ResultCode = &v
	return s
}
func (s *TaobaoTbkRelationRefundRpcResult) SetBizErrorDesc(v string) *TaobaoTbkRelationRefundRpcResult {
	s.BizErrorDesc = &v
	return s
}
func (s *TaobaoTbkRelationRefundRpcResult) SetBizErrorCode(v int64) *TaobaoTbkRelationRefundRpcResult {
	s.BizErrorCode = &v
	return s
}
