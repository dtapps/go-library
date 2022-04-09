package response

type TaobaoTopSdkFeedbackUploadResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   控制回传间隔（单位：秒）
	*/
	UploadInterval int64 `json:"upload_interval,omitempty" `
}
