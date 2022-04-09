package request

import (
	"topsdk/ability417/domain"
	"topsdk/util"
)

type TaobaoTbkDgPunishOrderGetRequest struct {
	/*
	   入参的对象     */
	AfOrderOption *domain.TaobaoTbkDgPunishOrderGetTopApiAfOrderOption `json:"af_order_option,omitempty" required:"false" `
}

func (s *TaobaoTbkDgPunishOrderGetRequest) SetAfOrderOption(v domain.TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) *TaobaoTbkDgPunishOrderGetRequest {
	s.AfOrderOption = &v
	return s
}

func (req *TaobaoTbkDgPunishOrderGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AfOrderOption != nil {
		paramMap["af_order_option"] = util.ConvertStruct(*req.AfOrderOption)
	}
	return paramMap
}

func (req *TaobaoTbkDgPunishOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
