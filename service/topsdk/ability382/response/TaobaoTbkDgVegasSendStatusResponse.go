package response

import (
	"topsdk/ability382/domain"
)

type TaobaoTbkDgVegasSendStatusResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果描述信息
	*/
	ResultMsg string `json:"result_msg,omitempty" `
	/*
	   返回结果封装对象
	*/
	Data domain.TaobaoTbkDgVegasSendStatusData `json:"data,omitempty" `
}
