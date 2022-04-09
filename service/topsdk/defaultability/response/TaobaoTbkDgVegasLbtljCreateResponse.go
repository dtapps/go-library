package response

import (
	"topsdk/defaultability/domain"
)

type TaobaoTbkDgVegasLbtljCreateResponse struct {

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
	Result domain.TaobaoTbkDgVegasLbtljCreateResult `json:"result,omitempty" `
}
