package response

type TaobaoTopAuthTokenRefreshResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回的是json信息
	*/
	TokenResult string `json:"token_result,omitempty" `
}
