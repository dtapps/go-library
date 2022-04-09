package request

type TaobaoTbkCouponGetRequest struct {
	/*
	   带券ID与商品ID的加密串     */
	Me *string `json:"me,omitempty" required:"false" `
	/*
	   商品ID     */
	ItemId *int64 `json:"item_id,omitempty" required:"false" `
	/*
	   券ID     */
	ActivityId *string `json:"activity_id,omitempty" required:"false" `
}

func (s *TaobaoTbkCouponGetRequest) SetMe(v string) *TaobaoTbkCouponGetRequest {
	s.Me = &v
	return s
}
func (s *TaobaoTbkCouponGetRequest) SetItemId(v int64) *TaobaoTbkCouponGetRequest {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkCouponGetRequest) SetActivityId(v string) *TaobaoTbkCouponGetRequest {
	s.ActivityId = &v
	return s
}

func (req *TaobaoTbkCouponGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Me != nil {
		paramMap["me"] = *req.Me
	}
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.ActivityId != nil {
		paramMap["activity_id"] = *req.ActivityId
	}
	return paramMap
}

func (req *TaobaoTbkCouponGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
