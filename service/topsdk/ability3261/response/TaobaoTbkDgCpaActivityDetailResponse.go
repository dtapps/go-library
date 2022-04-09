package response

import (
	"topsdk/ability3261/domain"
)

type TaobaoTbkDgCpaActivityDetailResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.TaobaoTbkDgCpaActivityDetailResult `json:"result,omitempty" `
}
