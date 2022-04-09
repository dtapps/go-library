package response

import (
	"topsdk/ability414/domain"
)

type TaobaoTbkOrderDetailsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   PublisherOrderDto
	*/
	Data domain.TaobaoTbkOrderDetailsGetOrderPage `json:"data,omitempty" `
}
