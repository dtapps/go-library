package response

import (
	"topsdk/ability1826/domain"
)

type TaobaoTbkDgVegasTljReportResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   结果
	*/
	Model domain.TaobaoTbkDgVegasTljReportInstanceDTO `json:"model,omitempty" `
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
