package response

import (
	"topsdk/ability376/domain"
)

type TaobaoTbkSpreadGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   传播形式对象列表
	*/
	Results []domain.TaobaoTbkSpreadGetTbkSpread `json:"results,omitempty" `
	/*
	   totalResults
	*/
	TotalResults int64 `json:"total_results,omitempty" `
}
