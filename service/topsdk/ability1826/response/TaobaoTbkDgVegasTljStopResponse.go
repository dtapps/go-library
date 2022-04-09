package response

import (
	"topsdk/ability1826/domain"
)

type TaobaoTbkDgVegasTljStopResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   model
	*/
	Model domain.TaobaoTbkDgVegasTljStopUpdateStatusResult `json:"model,omitempty" `
	/*
	   msgInfo
	*/
	MsgInfo string `json:"msg_info,omitempty" `
	/*
	   msgCode
	*/
	MsgCode string `json:"msg_code,omitempty" `
	/*
	   调用接口是否成功
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
}
