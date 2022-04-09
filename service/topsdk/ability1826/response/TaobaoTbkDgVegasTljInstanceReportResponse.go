package response

import (
	"topsdk/ability1826/domain"
)

type TaobaoTbkDgVegasTljInstanceReportResponse struct {

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
	Result domain.TaobaoTbkDgVegasTljInstanceReportResult `json:"result,omitempty" `
}
