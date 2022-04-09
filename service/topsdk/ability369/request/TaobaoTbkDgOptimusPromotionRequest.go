package request

type TaobaoTbkDgOptimusPromotionRequest struct {
	/*
	   页大小，一次请求请限制在10以内 defalutValue��10    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   第几页，默认：1 defalutValue��1    */
	PageNum *int64 `json:"page_num,omitempty" required:"false" `
	/*
	   mm_xxx_xxx_xxx的第3段数字     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！     */
	PromotionId *int64 `json:"promotion_id" required:"true" `
}

func (s *TaobaoTbkDgOptimusPromotionRequest) SetPageSize(v int64) *TaobaoTbkDgOptimusPromotionRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionRequest) SetPageNum(v int64) *TaobaoTbkDgOptimusPromotionRequest {
	s.PageNum = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionRequest) SetAdzoneId(v int64) *TaobaoTbkDgOptimusPromotionRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgOptimusPromotionRequest) SetPromotionId(v int64) *TaobaoTbkDgOptimusPromotionRequest {
	s.PromotionId = &v
	return s
}

func (req *TaobaoTbkDgOptimusPromotionRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNum != nil {
		paramMap["page_num"] = *req.PageNum
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.PromotionId != nil {
		paramMap["promotion_id"] = *req.PromotionId
	}
	return paramMap
}

func (req *TaobaoTbkDgOptimusPromotionRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
