package response

import (
	"topsdk/ability2474/domain"
)

type TaobaoTbkDgVegasSendReportResponse struct {

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
	Result domain.TaobaoTbkDgVegasSendReportResult `json:"result,omitempty" `
}
