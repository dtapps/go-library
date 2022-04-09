package request

import (
	"topsdk/ability373/domain"
	"topsdk/util"
)

type TaobaoJuItemsSearchRequest struct {
	/*
	   query     */
	ParamTopItemQuery *domain.TaobaoJuItemsSearchTopItemQuery `json:"param_top_item_query,omitempty" required:"false" `
}

func (s *TaobaoJuItemsSearchRequest) SetParamTopItemQuery(v domain.TaobaoJuItemsSearchTopItemQuery) *TaobaoJuItemsSearchRequest {
	s.ParamTopItemQuery = &v
	return s
}

func (req *TaobaoJuItemsSearchRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ParamTopItemQuery != nil {
		paramMap["param_top_item_query"] = util.ConvertStruct(*req.ParamTopItemQuery)
	}
	return paramMap
}

func (req *TaobaoJuItemsSearchRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
