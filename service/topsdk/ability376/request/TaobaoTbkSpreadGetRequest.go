package request

import (
	"topsdk/ability376/domain"
	"topsdk/util"
)

type TaobaoTbkSpreadGetRequest struct {
	/*
	   请求列表，内部包含多个url     */
	Requests *[]domain.TaobaoTbkSpreadGetTbkSpreadRequest `json:"requests" required:"true" `
}

func (s *TaobaoTbkSpreadGetRequest) SetRequests(v []domain.TaobaoTbkSpreadGetTbkSpreadRequest) *TaobaoTbkSpreadGetRequest {
	s.Requests = &v
	return s
}

func (req *TaobaoTbkSpreadGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Requests != nil {
		paramMap["requests"] = util.ConvertStructList(*req.Requests)
	}
	return paramMap
}

func (req *TaobaoTbkSpreadGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
