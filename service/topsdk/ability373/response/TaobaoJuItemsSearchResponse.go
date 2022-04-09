package response

import (
	"topsdk/ability373/domain"
)

type TaobaoJuItemsSearchResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Result domain.TaobaoJuItemsSearchPaginationResult `json:"result,omitempty" `
}
