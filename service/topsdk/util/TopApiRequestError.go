package util

import "fmt"

type TopApiRequestError struct {
	/*
	   System code
	*/
	TopCode int `json:"code,omitempty" `

	/*
	   System error message
	*/
	Msg string `json:"msg,omitempty" `

	/*
	   System sub code
	*/
	SubCode string `json:"sub_code,omitempty" `

	/*
	   System sub message
	*/
	SubMsg string `json:"sub_msg,omitempty" `

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `
}

func (e *TopApiRequestError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, sub_code: %s, sub_msg: %s ,request_id: %s", e.TopCode, e.Msg, e.SubCode, e.SubMsg, e.RequestId)
}
