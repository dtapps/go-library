package response

import (
	"topsdk/ability370/domain"
)

type TaobaoTbkDgMaterialOptionalResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   搜索到符合条件的结果总数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkDgMaterialOptionalMapData `json:"result_list,omitempty" `
	/*
	   本地化-lbs分页标识，请在下一次翻页时作为入参传入
	*/
	PageResultKey string `json:"page_result_key,omitempty" `
}
