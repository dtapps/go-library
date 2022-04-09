package response

import (
	"topsdk/ability417/domain"
)

type TaobaoTbkDgPunishOrderGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询的对象
	*/
	Result domain.TaobaoTbkDgPunishOrderGetRpcResult `json:"result,omitempty" `
}
