package domain

type TaobaoTbkDgCpaActivityDetailResult struct {
	/*
	   错误代码     */
	BizErrorFeature *string `json:"biz_error_feature,omitempty" `

	/*
	   返回素材id     */
	Data *TaobaoTbkDgCpaActivityDetailPageResult `json:"data,omitempty" `

	/*
	   是否成功     */
	Success *bool `json:"success,omitempty" `

	/*
	   结果码     */
	ResultCode *int64 `json:"result_code,omitempty" `

	/*
	   错误描述     */
	BizErrorDesc *string `json:"biz_error_desc,omitempty" `

	/*
	   错误代码     */
	BizErrorCode *int64 `json:"biz_error_code,omitempty" `

	/*
	   结果信息     */
	ResultMsg *string `json:"result_msg,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityDetailResult) SetBizErrorFeature(v string) *TaobaoTbkDgCpaActivityDetailResult {
	s.BizErrorFeature = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetData(v TaobaoTbkDgCpaActivityDetailPageResult) *TaobaoTbkDgCpaActivityDetailResult {
	s.Data = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetSuccess(v bool) *TaobaoTbkDgCpaActivityDetailResult {
	s.Success = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetResultCode(v int64) *TaobaoTbkDgCpaActivityDetailResult {
	s.ResultCode = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetBizErrorDesc(v string) *TaobaoTbkDgCpaActivityDetailResult {
	s.BizErrorDesc = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetBizErrorCode(v int64) *TaobaoTbkDgCpaActivityDetailResult {
	s.BizErrorCode = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResult) SetResultMsg(v string) *TaobaoTbkDgCpaActivityDetailResult {
	s.ResultMsg = &v
	return s
}
