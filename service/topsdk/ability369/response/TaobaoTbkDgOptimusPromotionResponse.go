package response

import (
	"topsdk/ability369/domain"
)

type TaobaoTbkDgOptimusPromotionResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkDgOptimusPromotionMapData `json:"result_list,omitempty" `
}
