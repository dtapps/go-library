package request

type TaobaoTbkDgOptimusMaterialRequest struct {
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   第几页，默认：1 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   mm_xxx_xxx_xxx的第三位     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)     */
	MaterialId *int64 `json:"material_id" required:"true" `
	/*
	   智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   智能匹配-设备号加密类型：MD5，类型为OAID时不传     */
	DeviceEncrypt *string `json:"device_encrypt,omitempty" required:"false" `
	/*
	   智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
	/*
	   内容专用-内容详情ID     */
	ContentId *int64 `json:"content_id,omitempty" required:"false" `
	/*
	   内容专用-内容渠道信息     */
	ContentSource *string `json:"content_source,omitempty" required:"false" `
	/*
	   商品ID，用于相似商品推荐     */
	ItemId *int64 `json:"item_id,omitempty" required:"false" `
	/*
	   选品库投放id     */
	FavoritesId *string `json:"favorites_id,omitempty" required:"false" `
}

func (s *TaobaoTbkDgOptimusMaterialRequest) SetPageSize(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetPageNo(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetAdzoneId(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetMaterialId(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.MaterialId = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetDeviceValue(v string) *TaobaoTbkDgOptimusMaterialRequest {
	s.DeviceValue = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetDeviceEncrypt(v string) *TaobaoTbkDgOptimusMaterialRequest {
	s.DeviceEncrypt = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetDeviceType(v string) *TaobaoTbkDgOptimusMaterialRequest {
	s.DeviceType = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetContentId(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.ContentId = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetContentSource(v string) *TaobaoTbkDgOptimusMaterialRequest {
	s.ContentSource = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetItemId(v int64) *TaobaoTbkDgOptimusMaterialRequest {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkDgOptimusMaterialRequest) SetFavoritesId(v string) *TaobaoTbkDgOptimusMaterialRequest {
	s.FavoritesId = &v
	return s
}

func (req *TaobaoTbkDgOptimusMaterialRequest) ToMap() map[string]interface{} {
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
	if req.MaterialId != nil {
		paramMap["material_id"] = *req.MaterialId
	}
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceEncrypt != nil {
		paramMap["device_encrypt"] = *req.DeviceEncrypt
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	if req.ContentId != nil {
		paramMap["content_id"] = *req.ContentId
	}
	if req.ContentSource != nil {
		paramMap["content_source"] = *req.ContentSource
	}
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.FavoritesId != nil {
		paramMap["favorites_id"] = *req.FavoritesId
	}
	return paramMap
}

func (req *TaobaoTbkDgOptimusMaterialRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
