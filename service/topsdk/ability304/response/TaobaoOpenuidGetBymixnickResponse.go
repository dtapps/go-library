package response

type TaobaoOpenuidGetBymixnickResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   OpenUID
	*/
	OpenUid string `json:"open_uid,omitempty" `
}
