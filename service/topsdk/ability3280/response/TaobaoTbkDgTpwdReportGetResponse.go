package response

import (
	"topsdk/ability3280/domain"
)

type TaobaoTbkDgTpwdReportGetResponse struct {

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
	Data domain.TaobaoTbkDgTpwdReportGetMapData `json:"data,omitempty" `
}
