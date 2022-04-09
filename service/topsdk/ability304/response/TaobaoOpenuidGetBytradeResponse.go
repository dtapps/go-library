package response

type TaobaoOpenuidGetBytradeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   当前交易tid对应买家的openuid
	*/
	OpenUid string `json:"open_uid,omitempty" `
}
